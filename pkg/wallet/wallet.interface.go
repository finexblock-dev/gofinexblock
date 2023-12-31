package wallet

import (
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet/structs"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Repository interface {
	types.Repository
	FindBlockchainByName(tx *gorm.DB, name string) (result *entity.Blockchain, err error)
	FindBlockchainByID(tx *gorm.DB, id uint) (result *entity.Blockchain, err error)

	FindCoinByID(tx *gorm.DB, id uint) (result *entity.Coin, err error)
	FindCoinByName(tx *gorm.DB, name string) (result *entity.Coin, err error)
	FindManyCoinByID(tx *gorm.DB, ids []uint) (result []*entity.Coin, err error)
	FindManyCoinByName(tx *gorm.DB, names []string) (result []*entity.Coin, err error)

	FindBlockNumberByCoinID(tx *gorm.DB, coinID uint) (result *entity.BlockNumber, err error)
	FindBlockNumberByID(tx *gorm.DB, id uint) (result *entity.BlockNumber, err error)
	UpdateBlockNumber(tx *gorm.DB, coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (result *entity.BlockNumber, err error)
	FindWalletByID(tx *gorm.DB, id uint) (result *entity.Wallet, err error)
	FindWalletByAddress(tx *gorm.DB, addr string) (result *entity.Wallet, err error)
	FindAllWallet(tx *gorm.DB, coinID uint) (result []*entity.Wallet, err error)

	ScanWalletByCoinID(tx *gorm.DB, coinID uint) (result []*entity.Wallet, err error)
	ScanWalletByUserID(tx *gorm.DB, userID uint) (result []*entity.Wallet, err error)
	ScanWalletByCond(tx *gorm.DB, userID, coinID uint) (result *entity.Wallet, err error)
	GetContractAddressByCoinID(tx *gorm.DB, coinID uint) (result *entity.SmartContract, err error)

	InsertWallet(tx *gorm.DB, wallet *entity.Wallet) (result *entity.Wallet, err error)
	UpdateWallet(tx *gorm.DB, id uint, address string) (result *entity.Wallet, err error)

	InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType entity.TransferType) (result *entity.CoinTransfer, err error)
	ScanCoinTransferByCond(tx *gorm.DB, userID, coinID uint, transferTypes []entity.TransferType, limit, offset int) (result []*entity.CoinTransfer, err error)

	FindCoinTransactionByID(tx *gorm.DB, id uint) (result *entity.CoinTransaction, err error)
	FindCoinTransactionByTxHash(tx *gorm.DB, txHash string) (result *entity.CoinTransaction, err error)
	ScanCoinTransactionByTransferID(tx *gorm.DB, transferID uint) (result []*entity.CoinTransaction, err error)
	ScanCoinTransactionByCond(tx *gorm.DB, transferID uint, status entity.TransactionStatus) (result []*entity.CoinTransaction, err error)
	InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus entity.TransactionStatus) (result *entity.CoinTransaction, err error)
	UpdateCoinTransactionHash(tx *gorm.DB, id uint, hash string) (result *entity.CoinTransaction, err error)
	UpdateCoinTransactionStatus(tx *gorm.DB, id uint, txStatus entity.TransactionStatus) (result *entity.CoinTransaction, err error)

	ScanWithdrawalRequestByUser(tx *gorm.DB, userID, coinID uint, limit, offset int) (result []*entity.WithdrawalRequest, err error)
	ScanWithdrawalRequestByStatus(tx *gorm.DB, status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error)
	ScanWithdrawalRequestByStatusWithLimitOffset(tx *gorm.DB, status entity.WithdrawalStatus, limit, offset int) (result []*entity.WithdrawalRequest, err error)
	ScanWithdrawalRequestByCond(tx *gorm.DB, coinID uint, status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error)
	ScanWithdrawalRequestByCondWithLimitOffset(tx *gorm.DB, coinID uint, status entity.WithdrawalStatus, limit, offset int) (result []*entity.WithdrawalRequest, err error)
	UpdateWithdrawalRequest(tx *gorm.DB, id uint, state entity.WithdrawalStatus) (result *entity.WithdrawalRequest, err error)
}

type Service interface {
	types.Service
	FindAllUserAssets(id uint) (result []*structs.Asset, err error)
	FindUserAssetsByCond(userID, coinID uint) (result *structs.Asset, err error)
	ScanCoinTransferByCond(userID, coinID uint, transferTypes []entity.TransferType, limit, offset int) (result []*entity.CoinTransfer, err error)

	FindBlockchainByName(name string) (result *entity.Blockchain, err error)
	FindBlockchainByID(id uint) (result *entity.Blockchain, err error)

	FindCoinByID(id uint) (result *entity.Coin, err error)
	FindCoinByName(name string) (result *entity.Coin, err error)
	FindManyCoinByID(ids []uint) (result []*entity.Coin, err error)
	FindManyCoinByName(names []string) (result []*entity.Coin, err error)

	FindBlockNumberByCoinID(coinID uint) (result *entity.BlockNumber, err error)
	FindBlockNumberByID(id uint) (result *entity.BlockNumber, err error)
	UpdateBlockNumber(coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (result *entity.BlockNumber, err error)

	FindWalletByID(id uint) (result *entity.Wallet, err error)
	FindWalletByAddress(addr string) (result *entity.Wallet, err error)
	FindAllWallet(coinID uint) (result []*entity.Wallet, err error)
	ScanWalletByCoinID(coinID uint) (result []*entity.Wallet, err error)
	ScanWalletByUserID(userID uint) (result []*entity.Wallet, err error)
	GetContractAddressByCoinID(coinID uint) (result *entity.SmartContract, err error)

	InsertWallet(wallet *entity.Wallet) (result *entity.Wallet, err error)
	UpdateWallet(id uint, address string) (result *entity.Wallet, err error)

	InsertCoinTransfer(walletID uint, amount decimal.Decimal, transferType entity.TransferType) (result *entity.CoinTransfer, err error)

	FindCoinTransactionByID(id uint) (result *entity.CoinTransaction, err error)
	FindCoinTransactionByTxHash(txHash string) (result *entity.CoinTransaction, err error)
	ScanCoinTransactionByTransferID(transferID uint) (result []*entity.CoinTransaction, err error)
	ScanCoinTransactionByCond(transferID uint, status entity.TransactionStatus) (result []*entity.CoinTransaction, err error)
	InsertCoinTransaction(transferID uint, txHash string, txStatus entity.TransactionStatus) (result *entity.CoinTransaction, err error)
	UpdateCoinTransactionHash(id uint, hash string) (result *entity.CoinTransaction, err error)
	UpdateCoinTransactionStatus(id uint, txStatus entity.TransactionStatus) (result *entity.CoinTransaction, err error)

	ScanWithdrawalRequestByUser(userID, coinID uint, limit, offset int) (result []*entity.WithdrawalRequest, err error)
	ScanWithdrawalRequestByStatus(status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error)
	ScanWithdrawalRequestByStatusWithLimitOffset(status entity.WithdrawalStatus, limit, offset int) (result []*entity.WithdrawalRequest, err error)

	ScanWithdrawalRequestByCond(coinID uint, status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error)
	ScanWithdrawalRequestByCondWithLimitOffset(coinID uint, status entity.WithdrawalStatus, limit, offset int) (result []*entity.WithdrawalRequest, err error)
	UpdateWithdrawalRequest(id uint, state entity.WithdrawalStatus) (result *entity.WithdrawalRequest, err error)

	BalanceUpdateInBatch(event []*grpc_order.BalanceUpdate) (err error)
}

func NewRepository(db *gorm.DB) Repository {
	return newWalletRepository(db)
}

func NewService(db *gorm.DB, cluster *redis.ClusterClient) Service {
	return newWalletService(db, cluster)
}