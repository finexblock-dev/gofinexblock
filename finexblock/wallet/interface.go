package wallet

import (
	context "context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Service interface {
	types.Service
	FindBlockchainByName(tx *gorm.DB, name string) (*wallet.Blockchain, error)
	FindBlockchainByID(tx *gorm.DB, id uint) (*wallet.Blockchain, error)

	FindCoinByID(tx *gorm.DB, id uint) (*wallet.Coin, error)
	FindCoinByName(tx *gorm.DB, name string) (*wallet.Coin, error)

	FindBlockNumberByCoinID(tx *gorm.DB, coinID uint) (*wallet.BlockNumber, error)
	FindBlockNumberByID(tx *gorm.DB, id uint) (*wallet.BlockNumber, error)
	UpdateBlockNumber(tx *gorm.DB, coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (*wallet.BlockNumber, error)

	FindWalletByID(tx *gorm.DB, id uint) (*wallet.Wallet, error)
	FindWalletByAddress(tx *gorm.DB, addr string) (*wallet.Wallet, error)
	FindAllWallet(tx *gorm.DB, coinID uint) ([]*wallet.Wallet, error)
	ScanWalletByCoinID(tx *gorm.DB, coinID uint) ([]*wallet.Wallet, error)
	ScanWalletByUserID(tx *gorm.DB, userID uint) ([]*wallet.Wallet, error)

	InsertWallet(tx *gorm.DB, wallet *wallet.Wallet) (*wallet.Wallet, error)
	UpdateWallet(tx *gorm.DB, id uint, address string) (*wallet.Wallet, error)

	InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType wallet.TransferType) (*wallet.CoinTransfer, error)

	FindCoinTransactionByTxHash(tx *gorm.DB, txHash string) (*wallet.CoinTransaction, error)
	ScanCoinTransactionByTransferID(tx *gorm.DB, transferID uint) ([]*wallet.CoinTransaction, error)
	ScanCoinTransactionByCond(tx *gorm.DB, transferID uint, status wallet.TransactionStatus) ([]*wallet.CoinTransaction, error)
	InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string) (*wallet.CoinTransaction, error)
	UpdateCoinTransaction(tx *gorm.DB, id uint, hash string) (*wallet.CoinTransaction, error)

	ScanWithdrawalRequestByStatus(tx *gorm.DB, status wallet.WithdrawalStatus) ([]*wallet.WithdrawalRequest, error)
	ScanWithdrawalRequestByCond(tx *gorm.DB, coinID uint, status wallet.WithdrawalStatus) ([]*wallet.WithdrawalRequest, error)
	UpdateWithdrawalRequest(tx *gorm.DB, id uint, state wallet.WithdrawalStatus) (*wallet.WithdrawalRequest, error)
}

type walletService struct {
	db *gorm.DB
}

func (w *walletService) Conn() *gorm.DB {
	return w.db
}

func (w *walletService) Ctx() context.Context {
	return context.Background()
}

func (w *walletService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func (w *walletService) Tx(level sql.IsolationLevel) *gorm.DB {
	return w.db.Begin(&sql.TxOptions{Isolation: level})
}

func newWalletService(db *gorm.DB) *walletService {
	return &walletService{db: db}
}

func NewService(db *gorm.DB) Service {
	return newWalletService(db)
}
