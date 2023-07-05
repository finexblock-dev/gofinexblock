package wallet

import (
	context "context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
)

type Service interface {
	types.Service
	FindBlockchainByName(name string) (*wallet.Blockchain, error)
	FindBlockchainByID(id uint) (*wallet.Blockchain, error)

	FindCoinByID(id uint) (*wallet.Coin, error)
	FindCoinByName(name string) (*wallet.Coin, error)

	FindBlockNumberByCoinID(tx *gorm.DB, coinID uint) (*wallet.BlockNumber, error)
	FindBlockNumberByID(tx *gorm.DB, id uint) (*wallet.BlockNumber, error)

	FindWalletByID(tx *gorm.DB, id uint) (*wallet.Wallet, error)
	FindWalletByAddress(tx *gorm.DB, addr string) (*wallet.Wallet, error)
	FindAllWallet(tx *gorm.DB, coinID uint) ([]*wallet.Wallet, error)
	ScanWalletByCoinID(tx *gorm.DB, coinID uint) ([]*wallet.Wallet, error)
	ScanWalletByUserID(tx *gorm.DB, userID uint) ([]*wallet.Wallet, error)

	InsertWallet(tx *gorm.DB, wallet *wallet.Wallet) (*wallet.Wallet, error)
	UpdateWallet(tx *gorm.DB, wallet *wallet.Wallet) (*wallet.Wallet, error)

	InsertCoinTransfer(tx *gorm.DB, coinTransfer *wallet.CoinTransfer) (*wallet.CoinTransfer, error)

	FindRecentCoinTransactionByTxHash(tx *gorm.DB, txHash string) (*wallet.CoinTransaction, error)
	InsertCoinTransaction(tx *gorm.DB, coinTransaction *wallet.CoinTransaction) (*wallet.CoinTransaction, error)

	ScanWithdrawalRequestByStatus(cond wallet.WithdrawalStatus) ([]*wallet.WithdrawalRequest, error)
	UpdateWithdrawalRequest(tx *gorm.DB, id uint, state wallet.WithdrawalStatus) error
}

type walletService struct {
	db *gorm.DB
}

func (w *walletService) Ctx() context.Context {
	//TODO implement me
	panic("implement me")
}

func (w *walletService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	//TODO implement me
	panic("implement me")
}

func (w *walletService) Tx(level sql.IsolationLevel) *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func newWalletService(db *gorm.DB) *walletService {
	return &walletService{db: db}
}

func NewService(db *gorm.DB) Service {
	return newWalletService(db)
}