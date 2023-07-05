package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) FindWalletByID(tx *gorm.DB, id uint) (*wallet.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletService) FindWalletByAddress(tx *gorm.DB, addr string) (*wallet.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletService) FindAllWallet(tx *gorm.DB, coinID uint) ([]*wallet.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletService) ScanWalletByCoinID(tx *gorm.DB, coinID uint) ([]*wallet.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletService) ScanWalletByUserID(tx *gorm.DB, userID uint) ([]*wallet.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletService) InsertWallet(tx *gorm.DB, wallet *wallet.Wallet) (*wallet.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletService) UpdateWallet(tx *gorm.DB, wallet *wallet.Wallet) (*wallet.Wallet, error) {
	//TODO implement me
	panic("implement me")
}