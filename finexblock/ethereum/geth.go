package ethereum

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/finexblock-dev/gofinexblock/finexblock/ethereum/hdwallet"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/big"
	"strconv"
)

type Interface interface {
	MasterWallet() *accounts.Account
	GetReceipt(ctx context.Context, txHash string) (*GetReceiptOutput, error)
	Transfer(ctx context.Context, userID, from, amount string) (string, error)
	CreateWallet(ctx context.Context, userID uint64) (string, error)
	GetBalance(ctx context.Context, address string) (string, error)
	GetBlockNumber(ctx context.Context) (uint64, error)
	GetBlocks(ctx context.Context, start, end uint64) ([][]byte, error)
	BlockNumber(c context.Context) (uint64, error)
	Nonce(c context.Context, account common.Address) (uint64, error)
	ChainID(c context.Context) (*big.Int, error)
	GasCap(c context.Context) (*big.Int, *big.Int, error)
	GasPrice(c context.Context) (*big.Int, error)
	BalanceAt(c context.Context, address common.Address) (*big.Int, error)
	SendRawTransaction(c context.Context, signedTx *types.Transaction) error
}

type GethClient struct {
	conn    *ethclient.Client
	master  *hdwallet.Wallet
	account *accounts.Account
}

func NewGethClient(rpcEndpoint, master string) (Interface, error) {
	var conn *ethclient.Client
	var err error
	// Dial rpc endpoint
	if conn, err = ethclient.Dial(rpcEndpoint); err != nil {
		return nil, err
	}

	wallet, err := HDWallet(master)
	if err != nil {
		log.Fatalf("failed to create hd wallet : %v", err)
	}

	path := DerivationPath("0", hdwallet.DefaultBaseDerivationPath.String())
	account, err := DerivedAccount(wallet, path)
	if err != nil {
		log.Fatalf("failed to derive account : %v", err)
	}

	return &GethClient{conn: conn, master: wallet, account: account}, nil
}

// MasterWallet returns master wallet
func (g *GethClient) MasterWallet() *accounts.Account {
	return g.account
}

// GetReceipt from transaction hash
func (g *GethClient) GetReceipt(ctx context.Context, txHash string) (*GetReceiptOutput, error) {
	hash := common.HexToHash(txHash)
	receipt, err := g.conn.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, err
	}
	return &GetReceiptOutput{
		TxHash:           receipt.TxHash.String(),
		Status:           receipt.Status,
		BlockHash:        receipt.BlockHash.String(),
		BlockNumber:      receipt.BlockNumber.String(),
		GasUsed:          receipt.GasUsed,
		TransactionIndex: uint64(receipt.TransactionIndex),
	}, nil
}

// Transfer When depositing, transfer from user account to hot wallet account
func (g *GethClient) Transfer(ctx context.Context, userID, from, amount string) (string, error) {

	// Master wallet address
	toAddress := g.MasterWallet().Address

	// User wallet address
	fromAddress := common.HexToAddress(from)

	// Convert string to decimal
	decimalAmount, err := decimal.NewFromString(amount)
	if err != nil {
		return "", fmt.Errorf("failed to convert string to decimal : %v", err)
	}

	value := big.NewInt(0)
	value.SetString(decimalAmount.String(), 10)

	// User custom nonce
	nonce, err := g.conn.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get user custom nonce : %v", err)
	}

	// Gas price
	gasTipCap, gasFeeCap, err := g.GasCap(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get gas cap : %v", err)
	}

	// Gas limit
	gasLimit := GasLimit()

	// Chain ID
	chainID, err := g.ChainID(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get chain id : %v", err)
	}

	log.Println("Gas Tip Cap : ", gasTipCap.String())
	log.Println("Gas Fee Cap : ", gasFeeCap.String())
	log.Println("Gas Limit : ", gasLimit)
	log.Println("Chain ID : ", chainID.String())
	log.Println("Nonce : ", nonce)
	log.Println("To Address : ", toAddress.String())
	log.Println("From Address : ", fromAddress.String())
	log.Println("Value : ", value.String())

	// Create transaction
	tx := NewTransaction(chainID, gasTipCap, gasFeeCap, value, nonce, gasLimit, toAddress)

	// Get user private key
	stringPrivateKey, err := PrivateKeyHex(g.master, userID)
	privateKey, err := crypto.HexToECDSA(stringPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to get user private key : %v", err)
	}

	log.Println("Signer Private key", stringPrivateKey)
	log.Println("Signer Public key", privateKey.Public())

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction : %v", err)
	}

	// Send transaction
	if err = g.conn.SendTransaction(ctx, signedTx); err != nil {
		return "", fmt.Errorf("failed to send transaction : %v", err)
	}

	return signedTx.Hash().Hex(), nil
}

// CreateWallet for user
func (g *GethClient) CreateWallet(ctx context.Context, userID uint64) (string, error) {
	path := DerivationPath(strconv.FormatUint(userID, 10), hdwallet.DefaultBaseDerivationPath.String())
	account, err := DerivedAccount(g.master, path)
	if err != nil {
		return "", status.Errorf(codes.Internal, "Error occurred while deriving account : %v", err)
	}
	return account.Address.String(), nil
}

// GetBalance of user
func (g *GethClient) GetBalance(ctx context.Context, address string) (string, error) {
	balance, err := g.conn.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return "", err
	}
	return balance.String(), nil
}

// GetBlockNumber of blockchain
func (g *GethClient) GetBlockNumber(ctx context.Context) (uint64, error) {
	blockNumber, err := g.conn.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}

// GetBlocks from blockchain
func (g *GethClient) GetBlocks(ctx context.Context, start, end uint64) ([][]byte, error) {
	if start > end {
		return nil, fmt.Errorf("invalid block range: start should be less than or equal to end")
	}

	var blockDataList [][]byte

	for i := start; i <= end; i++ {
		block, err := g.conn.BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		blockData, err := json.MarshalIndent(block.Transactions(), "", "  ")
		if err != nil {
			return nil, err
		}

		blockDataList = append(blockDataList, blockData)
	}

	return blockDataList, nil
}

// BlockNumber returns the current block number of the blockchain
func (g *GethClient) BlockNumber(c context.Context) (uint64, error) {
	return g.conn.BlockNumber(c)
}

// Nonce returns the nonce (transaction count) for an account
func (g *GethClient) Nonce(c context.Context, account common.Address) (uint64, error) {
	return g.conn.PendingNonceAt(c, account)
}

// ChainID returns the chain ID of the connected blockchain
func (g *GethClient) ChainID(c context.Context) (*big.Int, error) {
	return g.conn.NetworkID(c)
}

// GasCap returns the suggested gas tip cap and fee cap
func (g *GethClient) GasCap(c context.Context) (*big.Int, *big.Int, error) {
	tipCap, err := g.conn.SuggestGasTipCap(c)
	if err != nil {
		return nil, nil, err
	}
	feeCap, err := g.conn.SuggestGasPrice(c)
	if err != nil {
		return nil, nil, err
	}
	return tipCap, feeCap, err
}

func (g *GethClient) GasPrice(c context.Context) (*big.Int, error) {
	return g.conn.SuggestGasPrice(c)
}

func (g *GethClient) BalanceAt(c context.Context, address common.Address) (*big.Int, error) {
	return g.conn.BalanceAt(c, address, nil)
}

func (g *GethClient) SendRawTransaction(c context.Context, signedTx *types.Transaction) error {
	return g.conn.SendTransaction(c, signedTx)
}