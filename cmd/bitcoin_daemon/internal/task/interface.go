package bitcoin

import (
	"context"
	"github.com/finexblock-dev/gofinexblock/pkg/daemon"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/bitcoin"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type DepositDaemon interface {
	daemon.Daemon

	ListUnspentUTXO(ctx context.Context, address string) (*bitcoin.ListUnspentOutput, error)
	IsValidTransaction(tx *gorm.DB, txHash string) (bool, error)

	ScanWallets(tx *gorm.DB) ([]*entity.Wallet, error)
	InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus entity.TransactionStatus) (*entity.CoinTransaction, error)
	InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType entity.TransferType) (*entity.CoinTransfer, error)
	AcquireLock(uuid, coin string) (bool, error)
	ReleaseLock(uuid, coin string) error

	Gathering(ctx context.Context)
}

type WithdrawalDaemon interface {
	daemon.Daemon

	ScanWithdrawalRequests(tx *gorm.DB, status entity.WithdrawalStatus) ([]*entity.WithdrawalRequest, error)
	UpdateWithdrawalRequest(tx *gorm.DB, id uint, status entity.WithdrawalStatus) error
	Withdrawal(ctx context.Context, toAddress string, amount decimal.Decimal) (*bitcoin.SendToAddressOutput, error)
	ProcessPendingWithdrawalRequest() error
	ProcessApprovedWithdrawalRequest() error
	InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus entity.TransactionStatus) (*entity.CoinTransaction, error)
	InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType entity.TransferType) (*entity.CoinTransfer, error)
	UpdateCoinTransactionHash(tx *gorm.DB, id uint, txHash string) (*entity.CoinTransaction, error)
	UpdateCoinTransactionStatus(tx *gorm.DB, id uint, txStatus entity.TransactionStatus) (*entity.CoinTransaction, error)
}

func NewDeposit(proxy bitcoin.BitcoinProxyClient, cluster *redis.ClusterClient, db *gorm.DB, interval time.Duration, address string) DepositDaemon {
	return newDepositDaemon(proxy, cluster, db, interval, address)
}

func NewWithdrawal(proxy bitcoin.BitcoinProxyClient, cluster *redis.ClusterClient, db *gorm.DB, interval time.Duration) WithdrawalDaemon {
	return newWithdrawalDaemon(proxy, db, cluster, interval)
}