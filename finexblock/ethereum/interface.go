package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/finexblock-dev/gofinexblock/finexblock/ethereum/hdwallet"
	"log"
	"math/big"
)

type WalletService interface {
	MasterWallet() *accounts.Account
	CreateWallet(ctx context.Context, userID uint64) (string, error)
	BalanceAt(c context.Context, address common.Address) (*big.Int, error)
	GetBalance(ctx context.Context, address string) (string, error)
}

type BlockService interface {
	GetBlockNumber(ctx context.Context) (uint64, error)
	GetBlockTransactions(ctx context.Context, start, end uint64) ([]types.Transactions, error)
	BlockNumber(c context.Context) (uint64, error)
	ChainID(c context.Context) (*big.Int, error)
}

type TransactionService interface {
	GasCap(c context.Context) (*big.Int, *big.Int, error)
	GasPrice(c context.Context) (*big.Int, error)
	SendRawTransaction(c context.Context, signedTx *types.Transaction) error
	Nonce(c context.Context, account common.Address) (uint64, error)
	Transfer(ctx context.Context, userID, from, amount string) (string, error)
	GetReceipt(ctx context.Context, txHash string) (*types.Receipt, error)
}

type SmartContractService interface {
	ContractInstance(address common.Address)
}

type Service interface {
	WalletService
	BlockService
	TransactionService
}

type gethClient struct {
	conn    *ethclient.Client
	master  *hdwallet.Wallet
	account *accounts.Account
}

func newGethClient(rpcEndpoint, master string) (*gethClient, error) {
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

	return &gethClient{conn: conn, master: wallet, account: account}, nil
}

func NewService(rpcEndpoint, master string) (Service, error) {
	return newGethClient(rpcEndpoint, master)
}