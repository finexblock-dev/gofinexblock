package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) FindCoinByID(tx *gorm.DB, id uint) (*wallet.Coin, error) {
	var coin *wallet.Coin
	var err error

	if err = tx.Table(coin.TableName()).Where("id = ?", id).First(&coin).Error; err != nil {
		return nil, err
	}

	return coin, nil
}

func (w *walletService) FindCoinByName(tx *gorm.DB, name string) (*wallet.Coin, error) {
	var coin *wallet.Coin
	var err error

	if err = tx.Table(coin.TableName()).Where("name = ?", name).First(&coin).Error; err != nil {
		return nil, err
	}

	return coin, nil
}