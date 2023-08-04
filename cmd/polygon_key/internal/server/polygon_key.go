package server

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/finexblock-dev/gofinexblock/cmd/polygon_key/internal/config"
	geth "github.com/finexblock-dev/gofinexblock/pkg/ethereum"
	"github.com/finexblock-dev/gofinexblock/pkg/ethereum/hdwallet"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/polygon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/big"
	"net"
)

type PolygonKey struct {
	gethClient geth.Service
	master     *hdwallet.Wallet

	polygon.UnimplementedHDWalletServer
	polygon.UnimplementedTransactionServer
}

func NewPolygonKey(cfg *config.PolygonKeyConfiguration) (*PolygonKey, error) {
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
	return &PolygonKey{
		gethClient:                     conn,
		master:                         wallet,
		UnimplementedHDWalletServer:    polygon.UnimplementedHDWalletServer{},
		UnimplementedTransactionServer: polygon.UnimplementedTransactionServer{},
	}, nil
}

func (e *PolygonKey) Listen(gRPCServer *grpc.Server, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Error occurred while listening port on %v : %v", port, err)
	}
	log.Println("GRPC SERVER START")
	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatalf("Error occurred while serve listener... : %v", err)
	}
}

func (e *PolygonKey) Register(grpcServer *grpc.Server) {
	polygon.RegisterTransactionServer(grpcServer, e)
	polygon.RegisterHDWalletServer(grpcServer, e)
}

// CreateWallet for user
func (e *PolygonKey) CreateWallet(ctx context.Context, input *polygon.CreateWalletInput) (*polygon.CreateWalletOutput, error) {
	address, err := e.gethClient.CreateWallet(ctx, input.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occurred while create wallet: [%v]", err)
	}
	return &polygon.CreateWalletOutput{Address: address}, nil
}

// GetBalance of user
func (e *PolygonKey) GetBalance(ctx context.Context, input *polygon.GetBalanceInput) (*polygon.GetBalanceOutput, error) {
	balance, err := e.gethClient.GetBalance(ctx, input.Address)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occurred while get balance: [%v]", err)
	}
	return &polygon.GetBalanceOutput{Balance: balance}, nil
}

// GetReceipt from transaction hash
func (e *PolygonKey) GetReceipt(ctx context.Context, input *polygon.GetReceiptInput) (*polygon.GetReceiptOutput, error) {
	output, err := e.gethClient.GetReceipt(ctx, input.TxHash)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error occurred while get receipt: [%v]", err)
	}
	return &polygon.GetReceiptOutput{
		TxHash:           output.TxHash.String(),
		Status:           output.Status,
		BlockHash:        output.BlockHash.String(),
		BlockNumber:      output.BlockNumber.String(),
		GasUsed:          output.GasUsed,
		TransactionIndex: uint64(output.TransactionIndex),
	}, nil
}

// SendRawTransaction from hot wallet account when withdrawing(Send Raw Transaction)
func (e *PolygonKey) SendRawTransaction(ctx context.Context, input *polygon.SendRawTransactionInput) (*polygon.SendRawTransactionOutput, error) {
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

	return &polygon.SendRawTransactionOutput{
		Success: true,
		TxHash:  signedTx.Hash().Hex(),
	}, nil
}

// GetBlockNumber of blockchain
func (e *PolygonKey) GetBlockNumber(ctx context.Context, input *polygon.GetBlockNumberInput) (*polygon.GetBlockNumberOutput, error) {
	blockNumber, err := e.gethClient.GetBlockNumber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error occurred while get block number: [%v]", err)
	}
	return &polygon.GetBlockNumberOutput{BlockNumber: blockNumber}, nil
}

// GetBlocks from blockchain
func (e *PolygonKey) GetBlocks(ctx context.Context, input *polygon.GetBlocksInput) (*polygon.GetBlocksOutput, error) {
	var result []*polygon.TxData
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

			result = append(result, &polygon.TxData{
				ToAddress: tx.To().String(),
				Amount:    tx.Value().String(),
				TxHash:    tx.Hash().String(),
			})
		}
	}

	return &polygon.GetBlocksOutput{
		Result: result,
	}, nil
}

func (e *PolygonKey) mustEmbedUnimplementedTransactionServer() {
	//TODO implement me
	panic("implement me")
}

func (e *PolygonKey) mustEmbedUnimplementedWalletServer() {
	//TODO implement me
	panic("implement me")
}