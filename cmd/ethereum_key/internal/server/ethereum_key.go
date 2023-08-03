package server

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/finexblock-dev/gofinexblock/cmd/ethereum_key/internal/config"
	geth "github.com/finexblock-dev/gofinexblock/pkg/ethereum"
	"github.com/finexblock-dev/gofinexblock/pkg/ethereum/hdwallet"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/ethereum"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/big"
	"net"
)

type EthereumKey struct {
	gethClient geth.Service
	master     *hdwallet.Wallet

	ethereum.UnimplementedTransactionServer
	ethereum.UnimplementedHDWalletServer
}

func NewEthereumKey(cfg *config.EthereumKeyConfiguration) (*EthereumKey, error) {
	endpoint := fmt.Sprintf("https://%v/%v", cfg.Endpoint, cfg.Token)
	master := fmt.Sprintf("%v%v", cfg.First, cfg.Second)

	conn, err := geth.NewService(endpoint, master)
	if err != nil {
		return nil, err
	}

	// Get seed from mnemonic
	wallet, err := geth.HDWallet(master)
	if err != nil {
		return nil, err
	}
	return &EthereumKey{
		gethClient:                     conn,
		master:                         wallet,
		UnimplementedTransactionServer: ethereum.UnimplementedTransactionServer{},
		UnimplementedHDWalletServer:    ethereum.UnimplementedHDWalletServer{},
	}, nil
}

func (e *EthereumKey) Listen(gRPCServer *grpc.Server, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Error occurred while listening port on %v : %v", port, err)
	}
	log.Println("GRPC SERVER START")
	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatalf("Error occurred while serve listener... : %v", err)
	}
}

func (e *EthereumKey) Register(grpcServer *grpc.Server) {
	ethereum.RegisterTransactionServer(grpcServer, e)
	ethereum.RegisterHDWalletServer(grpcServer, e)
}

// CreateWallet for user
func (e *EthereumKey) CreateWallet(ctx context.Context, input *ethereum.CreateWalletInput) (*ethereum.CreateWalletOutput, error) {
	address, err := e.gethClient.CreateWallet(ctx, input.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occurred while create wallet: [%v]", err)
	}
	return &ethereum.CreateWalletOutput{Address: address}, nil
}

// GetBalance of user
func (e *EthereumKey) GetBalance(ctx context.Context, input *ethereum.GetBalanceInput) (*ethereum.GetBalanceOutput, error) {
	balance, err := e.gethClient.GetBalance(ctx, input.Address)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occurred while get balance: [%v]", err)
	}
	return &ethereum.GetBalanceOutput{Balance: balance}, nil
}

// GetReceipt from transaction hash
func (e *EthereumKey) GetReceipt(ctx context.Context, input *ethereum.GetReceiptInput) (*ethereum.GetReceiptOutput, error) {
	output, err := e.gethClient.GetReceipt(ctx, input.TxHash)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occurred while get receipt: [%v]", err)
	}
	return &ethereum.GetReceiptOutput{
		TxHash:           output.TxHash.String(),
		Status:           output.Status,
		BlockHash:        output.BlockHash.String(),
		BlockNumber:      output.BlockNumber.String(),
		GasUsed:          output.GasUsed,
		TransactionIndex: uint64(output.TransactionIndex),
	}, nil
}

// SendRawTransaction from hot wallet account when withdrawing(Send Raw Transaction)
func (e *EthereumKey) SendRawTransaction(ctx context.Context, input *ethereum.SendRawTransactionInput) (*ethereum.SendRawTransactionOutput, error) {
	var gasTipCap, gasFeeCap, chainID, value *big.Int
	var toAddress, fromAddress common.Address
	var privateKey *ecdsa.PrivateKey
	var stringPrivateKey string
	var gasLimit, nonce uint64
	var signedTx, tx *types.Transaction
	var err error

	toAddress = common.HexToAddress(input.To)
	fromAddress = common.HexToAddress(input.From)

	value = big.NewInt(0)
	value.SetString(input.Amount, 10)

	nonce, err = e.gethClient.Nonce(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	gasTipCap, gasFeeCap, err = e.gethClient.GasCap(ctx)
	if err != nil {
		return nil, err
	}

	gasLimit = geth.GasLimit()

	chainID, err = e.gethClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	tx = geth.NewTransaction(chainID, gasTipCap, gasFeeCap, value, nonce, gasLimit, toAddress)

	stringPrivateKey, err = geth.PrivateKeyHex(e.master, "0")
	privateKey, err = crypto.HexToECDSA(stringPrivateKey)
	if err != nil {
		return nil, err
	}

	signedTx, err = types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	if err = e.gethClient.SendRawTransaction(ctx, signedTx); err != nil {
		return nil, err
	}

	return &ethereum.SendRawTransactionOutput{
		Success: true,
		TxHash:  signedTx.Hash().Hex(),
	}, nil
}

// GetBlockNumber of blockchain
func (e *EthereumKey) GetBlockNumber(ctx context.Context, _ *ethereum.GetBlockNumberInput) (*ethereum.GetBlockNumberOutput, error) {
	blockNumber, err := e.gethClient.GetBlockNumber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occurred while get block number: [%v]", err)
	}
	return &ethereum.GetBlockNumberOutput{BlockNumber: blockNumber}, nil
}

// GetBlocks from blockchain
func (e *EthereumKey) GetBlocks(ctx context.Context, input *ethereum.GetBlocksInput) (*ethereum.GetBlocksOutput, error) {
	var result []*ethereum.TxData
	var blocks []types.Transactions
	var err error

	blocks, err = e.gethClient.GetBlockTransactions(ctx, input.Start, input.End)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occurred while get blocks: [%v]", err)
	}

	for _, block := range blocks {
		for _, tx := range block {
			// exception for contract creation
			if tx.To() == nil {
				continue
			}

			result = append(result, &ethereum.TxData{
				ToAddress: tx.To().String(),
				Amount:    tx.Value().String(),
				TxHash:    tx.Hash().String(),
			})
		}
	}

	return &ethereum.GetBlocksOutput{
		Result: result,
	}, nil
}

func (e *EthereumKey) mustEmbedUnimplementedTransactionServer() {
	//TODO implement me
	panic("implement me")
}

func (e *EthereumKey) mustEmbedUnimplementedHDWalletServer() {
	//TODO implement me
	panic("implement me")
}