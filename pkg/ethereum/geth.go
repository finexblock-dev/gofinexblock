package ethereum

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/finexblock-dev/gofinexblock/pkg/ethereum/hdwallet"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/big"
	"strconv"
)

// MasterWallet returns master wallet
func (g *gethClient) MasterWallet() *accounts.Account {
	return g.account
}

// GetReceipt from transaction hash
func (g *gethClient) GetReceipt(ctx context.Context, txHash string) (*types.Receipt, error) {
	hash := common.HexToHash(txHash)
	receipt, err := g.conn.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

// Transfer When depositing, transfer from user account to hot wallet account
func (g *gethClient) Transfer(ctx context.Context, userID, from, amount string) (string, error) {

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

	// Create transaction
	tx := NewTransaction(chainID, gasTipCap, gasFeeCap, value, nonce, gasLimit, toAddress)

	// Get user private key
	stringPrivateKey, err := PrivateKeyHex(g.master, userID)
	privateKey, err := crypto.HexToECDSA(stringPrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to get user private key : %v", err)
	}

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
func (g *gethClient) CreateWallet(ctx context.Context, userID uint64) (string, error) {
	path := DerivationPath(strconv.FormatUint(userID, 10), hdwallet.DefaultBaseDerivationPath.String())
	account, err := DerivedAccount(g.master, path)
	if err != nil {
		return "", status.Errorf(codes.Internal, "Error occurred while deriving account : %v", err)
	}
	return account.Address.String(), nil
}

// GetBalance of user
func (g *gethClient) GetBalance(ctx context.Context, address string) (string, error) {
	balance, err := g.conn.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return "", err
	}
	return balance.String(), nil
}

// GetBlockNumber of blockchain
func (g *gethClient) GetBlockNumber(ctx context.Context) (uint64, error) {
	blockNumber, err := g.conn.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}

// GetBlockTransactions from blockchain
func (g *gethClient) GetBlockTransactions(ctx context.Context, start, end uint64) ([]types.Transactions, error) {
	if start > end {
		return nil, fmt.Errorf("invalid block range: start should be less than or equal to end")
	}

	var blockDataList []types.Transactions

	for i := start; i <= end; i++ {
		block, err := g.conn.BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		blockDataList = append(blockDataList, g.Transactions(block))
	}

	return blockDataList, nil
}

func (g *gethClient) Transactions(block *types.Block) types.Transactions {
	return block.Transactions()
}

// BlockNumber returns the current block number of the blockchain
func (g *gethClient) BlockNumber(c context.Context) (uint64, error) {
	return g.conn.BlockNumber(c)
}

// Nonce returns the nonce (transaction count) for an account
func (g *gethClient) Nonce(c context.Context, account common.Address) (uint64, error) {
	return g.conn.PendingNonceAt(c, account)
}

// ChainID returns the chain ID of the connected blockchain
func (g *gethClient) ChainID(c context.Context) (*big.Int, error) {
	return g.conn.NetworkID(c)
}

// GasCap returns the suggested gas tip cap and fee cap
func (g *gethClient) GasCap(c context.Context) (*big.Int, *big.Int, error) {
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

func (g *gethClient) GasPrice(c context.Context) (*big.Int, error) {
	return g.conn.SuggestGasPrice(c)
}

func (g *gethClient) BalanceAt(c context.Context, address common.Address) (*big.Int, error) {
	return g.conn.BalanceAt(c, address, nil)
}

func (g *gethClient) SendRawTransaction(c context.Context, signedTx *types.Transaction) error {
	return g.conn.SendTransaction(c, signedTx)
}