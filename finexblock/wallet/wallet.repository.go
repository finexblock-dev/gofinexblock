package wallet

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type walletRepository struct {
	db *gorm.DB
}

func (w *walletRepository) FindWithdrawalRequestByID(tx *gorm.DB, id uint) (*wallet.WithdrawalRequest, error) {
	var _withdrawalRequest *wallet.WithdrawalRequest
	var err error

	if err = tx.Table(_withdrawalRequest.TableName()).Where("id = ?", id).First(&_withdrawalRequest).Error; err != nil {
		return nil, err
	}

	return _withdrawalRequest, nil
}

func (w *walletRepository) ScanWithdrawalRequestByStatus(tx *gorm.DB, status wallet.WithdrawalStatus) ([]*wallet.WithdrawalRequest, error) {
	var _withdrawalRequest []*wallet.WithdrawalRequest
	var table *wallet.WithdrawalRequest
	var err error

	if err = tx.Table(table.TableName()).Where("status = ?", status).Find(&_withdrawalRequest).Error; err != nil {
		return nil, err
	}

	return _withdrawalRequest, nil
}

func (w *walletRepository) ScanWithdrawalRequestByCond(tx *gorm.DB, coinID uint, status wallet.WithdrawalStatus) ([]*wallet.WithdrawalRequest, error) {
	var withdrawalRequests []*wallet.WithdrawalRequest
	var table *wallet.WithdrawalRequest

	err := tx.Table(table.TableName()).
		Select("withdrawal_request.*").
		Joins("join coin_transfer on coin_transfer.id = withdrawal_request.coin_transfer_id").
		Joins("join wallet on wallet.id = coin_transfer.wallet_id").
		Where("wallet.coin_id = ?", coinID).
		Where("withdrawal_request.status = ?", status).
		Find(&withdrawalRequests).Error

	if err != nil {
		return nil, err
	}

	return withdrawalRequests, nil
}

func (w *walletRepository) UpdateWithdrawalRequest(tx *gorm.DB, id uint, state wallet.WithdrawalStatus) (*wallet.WithdrawalRequest, error) {
	var _withdrawalRequest *wallet.WithdrawalRequest
	var err error

	if err = tx.Table(_withdrawalRequest.TableName()).Where("id = ?", id).Update("status", state).Error; err != nil {
		return nil, err
	}

	return w.FindWithdrawalRequestByID(tx, id)
}

func (w *walletRepository) FindWalletByID(tx *gorm.DB, id uint) (*wallet.Wallet, error) {
	var _wallet *wallet.Wallet
	var err error

	if err = tx.Table(_wallet.TableName()).Where("id = ?", id).First(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) FindWalletByAddress(tx *gorm.DB, addr string) (*wallet.Wallet, error) {
	var _wallet *wallet.Wallet
	var err error

	if err = tx.Table(_wallet.TableName()).Where("address = ?", addr).First(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) FindAllWallet(tx *gorm.DB, coinID uint) ([]*wallet.Wallet, error) {
	var _wallet []*wallet.Wallet
	var _table *wallet.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("coin_id = ?", coinID).Find(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) ScanWalletByCoinID(tx *gorm.DB, coinID uint) ([]*wallet.Wallet, error) {
	var _wallet []*wallet.Wallet
	var _table *wallet.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("coin_id = ?", coinID).Find(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) ScanWalletByUserID(tx *gorm.DB, userID uint) ([]*wallet.Wallet, error) {
	var _wallet []*wallet.Wallet
	var _table *wallet.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("user_id = ?", userID).Find(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) InsertWallet(tx *gorm.DB, _wallet *wallet.Wallet) (*wallet.Wallet, error) {
	var err error

	if err = tx.Table(_wallet.TableName()).Create(_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) UpdateWallet(tx *gorm.DB, id uint, address string) (*wallet.Wallet, error) {
	var _table *wallet.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("id = ?", id).Update("address", address).Error; err != nil {
		return nil, err
	}

	return w.FindWalletByID(tx, id)
}

func (w *walletRepository) FindSmartContractByCoinID(tx *gorm.DB, coinID uint) (*wallet.SmartContract, error) {
	var smartContract *wallet.SmartContract
	var err error

	if err = tx.Table(smartContract.TableName()).Where("coin_id = ?", coinID).First(&smartContract).Error; err != nil {
		return nil, err
	}

	return smartContract, nil
}

func (w *walletRepository) FindSmartContractByID(tx *gorm.DB, id uint) (*wallet.SmartContract, error) {
	var smartContract *wallet.SmartContract
	var err error

	if err = tx.Table(smartContract.TableName()).Where("id = ?", id).First(&smartContract).Error; err != nil {
		return nil, err
	}

	return smartContract, nil
}

func (w *walletRepository) InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType wallet.TransferType) (*wallet.CoinTransfer, error) {
	var coinTransfer *wallet.CoinTransfer
	var err error

	coinTransfer = &wallet.CoinTransfer{
		WalletID:     walletID,
		Amount:       amount,
		TransferType: transferType,
	}

	if err = tx.Table(coinTransfer.TableName()).Create(coinTransfer).Error; err != nil {
		return nil, err
	}

	return coinTransfer, nil
}

func (w *walletRepository) ScanCoinTransactionByTransferID(tx *gorm.DB, transferID uint) ([]*wallet.CoinTransaction, error) {
	var coinTransaction []*wallet.CoinTransaction
	var _table *wallet.CoinTransaction
	var err error

	if err = tx.Table(_table.TableName()).Where("coin_transfer_id = ?", transferID).Find(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletRepository) InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus wallet.TransactionStatus) (*wallet.CoinTransaction, error) {
	var err error
	var coinTransaction *wallet.CoinTransaction

	coinTransaction = &wallet.CoinTransaction{
		CoinTransferID: transferID,
		TxHash:         txHash,
		Status:         txStatus,
	}

	if err = tx.Table(coinTransaction.TableName()).Create(coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletRepository) FindCoinTransactionByTxHash(tx *gorm.DB, txHash string) (*wallet.CoinTransaction, error) {
	var coinTransaction *wallet.CoinTransaction
	var err error

	if err = tx.Table(coinTransaction.TableName()).Where("tx_hash = ?", txHash).First(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletRepository) FindCoinTransactionByID(tx *gorm.DB, id uint) (*wallet.CoinTransaction, error) {
	var coinTransaction *wallet.CoinTransaction
	var err error

	if err = tx.Table(coinTransaction.TableName()).Where("id = ?", id).First(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletRepository) ScanCoinTransactionByCond(tx *gorm.DB, transferID uint, status wallet.TransactionStatus) ([]*wallet.CoinTransaction, error) {
	var coinTransaction []*wallet.CoinTransaction
	var table *wallet.CoinTransaction
	var err error

	if err = tx.Table(table.TableName()).Where("coin_transfer_id = ?", transferID).Where("status = ?", status).Find(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletRepository) UpdateCoinTransactionHash(tx *gorm.DB, id uint, hash string) (*wallet.CoinTransaction, error) {
	var table *wallet.CoinTransaction
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("tx_hash", hash).Error; err != nil {
		return nil, err
	}

	return w.FindCoinTransactionByTxHash(tx, hash)
}

func (w *walletRepository) UpdateCoinTransactionStatus(tx *gorm.DB, id uint, txStatus wallet.TransactionStatus) (*wallet.CoinTransaction, error) {
	var table *wallet.CoinTransaction
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("status", txStatus).Error; err != nil {
		return nil, err
	}

	return w.FindCoinTransactionByID(tx, id)
}

func (w *walletRepository) FindCoinByID(tx *gorm.DB, id uint) (*wallet.Coin, error) {
	var coin *wallet.Coin
	var err error

	if err = tx.Table(coin.TableName()).Where("id = ?", id).First(&coin).Error; err != nil {
		return nil, err
	}

	return coin, nil
}

func (w *walletRepository) FindCoinByName(tx *gorm.DB, name string) (*wallet.Coin, error) {
	var coin *wallet.Coin
	var err error

	if err = tx.Table(coin.TableName()).Where("name = ?", name).First(&coin).Error; err != nil {
		return nil, err
	}

	return coin, nil
}

func (w *walletRepository) FindBlockchainByName(tx *gorm.DB, name string) (*wallet.Blockchain, error) {
	var blockchain *wallet.Blockchain
	var err error

	if err = tx.Table(blockchain.TableName()).Where("name = ?", name).First(&blockchain).Error; err != nil {
		return nil, err
	}

	return blockchain, nil
}

func (w *walletRepository) FindBlockchainByID(tx *gorm.DB, id uint) (*wallet.Blockchain, error) {
	var blockchain *wallet.Blockchain
	var err error

	if err = tx.Table(blockchain.TableName()).Where("id = ?", id).First(&blockchain).Error; err != nil {
		return nil, err
	}

	return blockchain, nil
}

func (w *walletRepository) FindBlockNumberByCoinID(tx *gorm.DB, coinID uint) (*wallet.BlockNumber, error) {
	var blockNumber *wallet.BlockNumber
	var err error

	if err = tx.Table(blockNumber.TableName()).Where("coin_id = ?", coinID).First(&blockNumber).Error; err != nil {
		return nil, err
	}

	return blockNumber, nil
}

func (w *walletRepository) FindBlockNumberByID(tx *gorm.DB, id uint) (*wallet.BlockNumber, error) {
	var blockNumber *wallet.BlockNumber
	var err error

	if err = tx.Table(blockNumber.TableName()).Where("id = ?", id).First(&blockNumber).Error; err != nil {
		return nil, err
	}

	return blockNumber, nil
}

func (w *walletRepository) UpdateBlockNumber(tx *gorm.DB, coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (*wallet.BlockNumber, error) {
	var err error
	var table *wallet.BlockNumber

	if err = tx.Table(table.TableName()).Where("coin_id = ?", coinID).Update("from_block", fromBlockNumber).Update("to_block", toBlockNumber).Error; err != nil {
		return nil, err
	}

	return w.FindBlockNumberByCoinID(tx, coinID)
}

func (w *walletRepository) Conn() *gorm.DB {
	return w.db
}

func (w *walletRepository) Ctx() context.Context {
	return context.Background()
}

func (w *walletRepository) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func (w *walletRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return w.db.Begin(&sql.TxOptions{Isolation: level})
}

func newWalletRepository(db *gorm.DB) *walletRepository {
	return &walletRepository{db: db}
}
