package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Repository interface {
	types.Repository
	FindBlockchainByName(tx *gorm.DB, name string) (result *wallet.Blockchain, err error)
	FindBlockchainByID(tx *gorm.DB, id uint) (result *wallet.Blockchain, err error)

	FindCoinByID(tx *gorm.DB, id uint) (result *wallet.Coin, err error)
	FindCoinByName(tx *gorm.DB, name string) (result *wallet.Coin, err error)

	FindBlockNumberByCoinID(tx *gorm.DB, coinID uint) (result *wallet.BlockNumber, err error)
	FindBlockNumberByID(tx *gorm.DB, id uint) (result *wallet.BlockNumber, err error)
	UpdateBlockNumber(tx *gorm.DB, coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (result *wallet.BlockNumber, err error)

	FindWalletByID(tx *gorm.DB, id uint) (result *wallet.Wallet, err error)
	FindWalletByAddress(tx *gorm.DB, addr string) (result *wallet.Wallet, err error)
	FindAllWallet(tx *gorm.DB, coinID uint) (result []*wallet.Wallet, err error)
	ScanWalletByCoinID(tx *gorm.DB, coinID uint) (result []*wallet.Wallet, err error)
	ScanWalletByUserID(tx *gorm.DB, userID uint) (result []*wallet.Wallet, err error)

	InsertWallet(tx *gorm.DB, wallet *wallet.Wallet) (result *wallet.Wallet, err error)
	UpdateWallet(tx *gorm.DB, id uint, address string) (result *wallet.Wallet, err error)

	InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType wallet.TransferType) (result *wallet.CoinTransfer, err error)

	FindCoinTransactionByID(tx *gorm.DB, id uint) (result *wallet.CoinTransaction, err error)
	FindCoinTransactionByTxHash(tx *gorm.DB, txHash string) (result *wallet.CoinTransaction, err error)
	ScanCoinTransactionByTransferID(tx *gorm.DB, transferID uint) (result []*wallet.CoinTransaction, err error)
	ScanCoinTransactionByCond(tx *gorm.DB, transferID uint, status wallet.TransactionStatus) (result []*wallet.CoinTransaction, err error)
	InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus wallet.TransactionStatus) (result *wallet.CoinTransaction, err error)
	UpdateCoinTransactionHash(tx *gorm.DB, id uint, hash string) (result *wallet.CoinTransaction, err error)
	UpdateCoinTransactionStatus(tx *gorm.DB, id uint, txStatus wallet.TransactionStatus) (result *wallet.CoinTransaction, err error)

	ScanWithdrawalRequestByStatus(tx *gorm.DB, status wallet.WithdrawalStatus) (result []*wallet.WithdrawalRequest, err error)
	ScanWithdrawalRequestByCond(tx *gorm.DB, coinID uint, status wallet.WithdrawalStatus) (result []*wallet.WithdrawalRequest, err error)
	UpdateWithdrawalRequest(tx *gorm.DB, id uint, state wallet.WithdrawalStatus) (result *wallet.WithdrawalRequest, err error)
}

type Service interface {
	types.Service
	FindBlockchainByName(name string) (result *wallet.Blockchain, err error)
	FindBlockchainByID(id uint) (result *wallet.Blockchain, err error)

	FindCoinByID(id uint) (result *wallet.Coin, err error)
	FindCoinByName(name string) (result *wallet.Coin, err error)

	FindBlockNumberByCoinID(coinID uint) (result *wallet.BlockNumber, err error)
	FindBlockNumberByID(id uint) (result *wallet.BlockNumber, err error)
	UpdateBlockNumber(coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (result *wallet.BlockNumber, err error)

	FindWalletByID(id uint) (result *wallet.Wallet, err error)
	FindWalletByAddress(addr string) (result *wallet.Wallet, err error)
	FindAllWallet(coinID uint) (result []*wallet.Wallet, err error)
	ScanWalletByCoinID(coinID uint) (result []*wallet.Wallet, err error)
	ScanWalletByUserID(userID uint) (result []*wallet.Wallet, err error)

	InsertWallet(wallet *wallet.Wallet) (result *wallet.Wallet, err error)
	UpdateWallet(id uint, address string) (result *wallet.Wallet, err error)

	InsertCoinTransfer(walletID uint, amount decimal.Decimal, transferType wallet.TransferType) (result *wallet.CoinTransfer, err error)

	FindCoinTransactionByID(id uint) (result *wallet.CoinTransaction, err error)
	FindCoinTransactionByTxHash(txHash string) (result *wallet.CoinTransaction, err error)
	ScanCoinTransactionByTransferID(transferID uint) (result []*wallet.CoinTransaction, err error)
	ScanCoinTransactionByCond(transferID uint, status wallet.TransactionStatus) (result []*wallet.CoinTransaction, err error)
	InsertCoinTransaction(transferID uint, txHash string, txStatus wallet.TransactionStatus) (result *wallet.CoinTransaction, err error)
	UpdateCoinTransactionHash(id uint, hash string) (result *wallet.CoinTransaction, err error)
	UpdateCoinTransactionStatus(id uint, txStatus wallet.TransactionStatus) (result *wallet.CoinTransaction, err error)

	ScanWithdrawalRequestByStatus(status wallet.WithdrawalStatus) (result []*wallet.WithdrawalRequest, err error)
	ScanWithdrawalRequestByCond(coinID uint, status wallet.WithdrawalStatus) (result []*wallet.WithdrawalRequest, err error)
	UpdateWithdrawalRequest(id uint, state wallet.WithdrawalStatus) (result *wallet.WithdrawalRequest, err error)
}

func NewRepository(db *gorm.DB) Repository {
	return newWalletRepository(db)
}

func NewService(db *gorm.DB) Service {
	return newWalletService(NewRepository(db))
}
