package polygon

import (
	"context"
	"github.com/finexblock-dev/gofinexblock/finexblock/daemon"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/polygon"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type DepositDaemon interface {
	daemon.Daemon

	InsertErrLog(err error) error
	GetCurrentBlockNumber(ctx context.Context) (uint64, error)
	GetBlock(ctx context.Context, start, end decimal.Decimal) ([]*polygon.TxData, error)
	GetBalance(ctx context.Context, address string) (*polygon.GetBalanceOutput, error)
	IsValidTransaction(tx *gorm.DB, txHash string) (bool, error)
	ScanWallets(tx *gorm.DB) ([]*entity.Wallet, error)
	InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus entity.TransactionStatus) (*entity.CoinTransaction, error)
	InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType entity.TransferType) (*entity.CoinTransfer, error)
	AcquireLock(uuid, coin string) (bool, error)
	Transfer(ctx context.Context, from, to string, amount decimal.Decimal) (*polygon.SendRawTransactionOutput, error)
	CheckReceipt(txHash string) (*polygon.GetReceiptOutput, error)
	ReleaseLock(uuid, coin string) error
	Gathering(ctx context.Context)
}

type WithdrawalDaemon interface {
	daemon.Daemon

	GetReceipt(ctx context.Context, txHash string) (*polygon.GetReceiptOutput, error)

	ScanWithdrawalRequests(tx *gorm.DB, status entity.WithdrawalStatus) ([]*entity.WithdrawalRequest, error)
	UpdateWithdrawalRequest(tx *gorm.DB, id uint, status entity.WithdrawalStatus) error
	Transfer(ctx context.Context, from, to string, amount decimal.Decimal) (*polygon.SendRawTransactionOutput, error)
	ProcessPendingWithdrawalRequest() error
	ProcessApprovedWithdrawalRequest() error
	InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus entity.TransactionStatus) (*entity.CoinTransaction, error)
	InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType entity.TransferType) (*entity.CoinTransfer, error)
	UpdateCoinTransactionHash(tx *gorm.DB, id uint, txHash string) (*entity.CoinTransaction, error)
	UpdateCoinTransactionStatus(tx *gorm.DB, id uint, txStatus entity.TransactionStatus) (*entity.CoinTransaction, error)
}

func NewDeposit(proxy polygon.PolygonProxyClient, centralWalletAddress string, db *gorm.DB, client *redis.ClusterClient, interval time.Duration) DepositDaemon {
	return newDepositDaemon(proxy, centralWalletAddress, db, client, interval)
}

func NewWithdrawal(proxy polygon.PolygonProxyClient, centralWalletAddress string, db *gorm.DB, client *redis.ClusterClient, interval time.Duration) WithdrawalDaemon {
	return newWithdrawalDaemon(proxy, centralWalletAddress, db, client, interval)
}
