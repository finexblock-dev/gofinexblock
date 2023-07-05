package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func (w *walletService) InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType wallet.TransferType) (*wallet.CoinTransfer, error) {
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