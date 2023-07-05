package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) FindBlockNumberByCoinID(tx *gorm.DB, coinID uint) (*wallet.BlockNumber, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletService) FindBlockNumberByID(tx *gorm.DB, id uint) (*wallet.BlockNumber, error) {
	//TODO implement me
	panic("implement me")
}