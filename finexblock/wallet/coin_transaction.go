package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) FindRecentCoinTransactionByTxHash(tx *gorm.DB, txHash string) (*wallet.CoinTransaction, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletService) InsertCoinTransaction(tx *gorm.DB, coinTransaction *wallet.CoinTransaction) (*wallet.CoinTransaction, error) {
	//TODO implement me
	panic("implement me")
}