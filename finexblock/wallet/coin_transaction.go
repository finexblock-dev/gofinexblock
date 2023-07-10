package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) ScanCoinTransactionByTransferID(tx *gorm.DB, transferID uint) ([]*wallet.CoinTransaction, error) {
	var coinTransaction []*wallet.CoinTransaction
	var _table *wallet.CoinTransaction
	var err error

	if err = tx.Table(_table.TableName()).Where("coin_transfer_id = ?", transferID).Find(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletService) InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus wallet.TransactionStatus) (*wallet.CoinTransaction, error) {
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

func (w *walletService) FindCoinTransactionByTxHash(tx *gorm.DB, txHash string) (*wallet.CoinTransaction, error) {
	var coinTransaction *wallet.CoinTransaction
	var err error

	if err = tx.Table(coinTransaction.TableName()).Where("tx_hash = ?", txHash).First(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletService) FindCoinTransactionByID(tx *gorm.DB, id uint) (*wallet.CoinTransaction, error) {
	var coinTransaction *wallet.CoinTransaction
	var err error

	if err = tx.Table(coinTransaction.TableName()).Where("id = ?", id).First(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletService) ScanCoinTransactionByCond(tx *gorm.DB, transferID uint, status wallet.TransactionStatus) ([]*wallet.CoinTransaction, error) {
	var coinTransaction []*wallet.CoinTransaction
	var table *wallet.CoinTransaction
	var err error

	if err = tx.Table(table.TableName()).Where("coin_transfer_id = ?", transferID).Where("status = ?", status).Find(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletService) UpdateCoinTransactionHash(tx *gorm.DB, id uint, hash string) (*wallet.CoinTransaction, error) {
	var table *wallet.CoinTransaction
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("tx_hash", hash).Error; err != nil {
		return nil, err
	}

	return w.FindCoinTransactionByTxHash(tx, hash)
}

func (w *walletService) UpdateCoinTransactionStatus(tx *gorm.DB, id uint, txStatus wallet.TransactionStatus) (*wallet.CoinTransaction, error) {
	var table *wallet.CoinTransaction
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("status", txStatus).Error; err != nil {
		return nil, err
	}

	return w.FindCoinTransactionByID(tx, id)
}