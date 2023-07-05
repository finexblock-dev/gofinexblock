package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) InsertCoinTransfer(tx *gorm.DB, coinTransfer *wallet.CoinTransfer) (*wallet.CoinTransfer, error) {
	//TODO implement me
	panic("implement me")
}