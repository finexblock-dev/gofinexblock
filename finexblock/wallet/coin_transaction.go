package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) ScanCoinTransactionByTransferID(tx *gorm.DB, txHash string) ([]*wallet.CoinTransaction, error) {
	var coinTransaction []*wallet.CoinTransaction
	var _table *wallet.CoinTransaction
	var err error

	if err = tx.Table(_table.TableName()).Where("tx_hash = ?", txHash).Find(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletService) InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string) (*wallet.CoinTransaction, error) {
	var err error
	var coinTransaction *wallet.CoinTransaction

	coinTransaction = &wallet.CoinTransaction{
		CoinTransferID: transferID,
		TxHash:         txHash,
	}

	if err = tx.Table(coinTransaction.TableName()).Create(coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}