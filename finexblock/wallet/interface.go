package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Repository interface {
	types.Repository
	FindBlockchainByName(tx *gorm.DB, name string) (result *entity.Blockchain, err error)
	FindBlockchainByID(tx *gorm.DB, id uint) (result *entity.Blockchain, err error)

	FindCoinByID(tx *gorm.DB, id uint) (result *entity.Coin, err error)
	FindCoinByName(tx *gorm.DB, name string) (result *entity.Coin, err error)

	FindBlockNumberByCoinID(tx *gorm.DB, coinID uint) (result *entity.BlockNumber, err error)
	FindBlockNumberByID(tx *gorm.DB, id uint) (result *entity.BlockNumber, err error)
	UpdateBlockNumber(tx *gorm.DB, coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (result *entity.BlockNumber, err error)

	FindWalletByID(tx *gorm.DB, id uint) (result *entity.Wallet, err error)
	FindWalletByAddress(tx *gorm.DB, addr string) (result *entity.Wallet, err error)
	FindAllWallet(tx *gorm.DB, coinID uint) (result []*entity.Wallet, err error)
	ScanWalletByCoinID(tx *gorm.DB, coinID uint) (result []*entity.Wallet, err error)
	ScanWalletByUserID(tx *gorm.DB, userID uint) (result []*entity.Wallet, err error)

	InsertWallet(tx *gorm.DB, wallet *entity.Wallet) (result *entity.Wallet, err error)
	UpdateWallet(tx *gorm.DB, id uint, address string) (result *entity.Wallet, err error)

	InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType entity.TransferType) (result *entity.CoinTransfer, err error)

	FindCoinTransactionByID(tx *gorm.DB, id uint) (result *entity.CoinTransaction, err error)
	FindCoinTransactionByTxHash(tx *gorm.DB, txHash string) (result *entity.CoinTransaction, err error)
	ScanCoinTransactionByTransferID(tx *gorm.DB, transferID uint) (result []*entity.CoinTransaction, err error)
	ScanCoinTransactionByCond(tx *gorm.DB, transferID uint, status entity.TransactionStatus) (result []*entity.CoinTransaction, err error)
	InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus entity.TransactionStatus) (result *entity.CoinTransaction, err error)
	UpdateCoinTransactionHash(tx *gorm.DB, id uint, hash string) (result *entity.CoinTransaction, err error)
	UpdateCoinTransactionStatus(tx *gorm.DB, id uint, txStatus entity.TransactionStatus) (result *entity.CoinTransaction, err error)

	ScanWithdrawalRequestByStatus(tx *gorm.DB, status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error)
	ScanWithdrawalRequestByCond(tx *gorm.DB, coinID uint, status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error)
	UpdateWithdrawalRequest(tx *gorm.DB, id uint, state entity.WithdrawalStatus) (result *entity.WithdrawalRequest, err error)
}

type Service interface {
	types.Service
	FindBlockchainByName(name string) (result *entity.Blockchain, err error)
	FindBlockchainByID(id uint) (result *entity.Blockchain, err error)

	FindCoinByID(id uint) (result *entity.Coin, err error)
	FindCoinByName(name string) (result *entity.Coin, err error)

	FindBlockNumberByCoinID(coinID uint) (result *entity.BlockNumber, err error)
	FindBlockNumberByID(id uint) (result *entity.BlockNumber, err error)
	UpdateBlockNumber(coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (result *entity.BlockNumber, err error)

	FindWalletByID(id uint) (result *entity.Wallet, err error)
	FindWalletByAddress(addr string) (result *entity.Wallet, err error)
	FindAllWallet(coinID uint) (result []*entity.Wallet, err error)
	ScanWalletByCoinID(coinID uint) (result []*entity.Wallet, err error)
	ScanWalletByUserID(userID uint) (result []*entity.Wallet, err error)

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

	ScanWithdrawalRequestByStatus(status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error)
	ScanWithdrawalRequestByCond(coinID uint, status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error)
	UpdateWithdrawalRequest(id uint, state entity.WithdrawalStatus) (result *entity.WithdrawalRequest, err error)
}

func NewRepository(db *gorm.DB) Repository {
	return newWalletRepository(db)
}

func NewService(db *gorm.DB) Service {
	return newWalletService(NewRepository(db))
}
