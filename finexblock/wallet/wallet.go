package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) FindWalletByID(tx *gorm.DB, id uint) (*wallet.Wallet, error) {
	var _wallet *wallet.Wallet
	var err error

	if err = tx.Table(_wallet.TableName()).Where("id = ?", id).First(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletService) FindWalletByAddress(tx *gorm.DB, addr string) (*wallet.Wallet, error) {
	var _wallet *wallet.Wallet
	var err error

	if err = tx.Table(_wallet.TableName()).Where("address = ?", addr).First(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletService) FindAllWallet(tx *gorm.DB, coinID uint) ([]*wallet.Wallet, error) {
	var _wallet []*wallet.Wallet
	var _table *wallet.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("coin_id = ?", coinID).Find(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletService) ScanWalletByCoinID(tx *gorm.DB, coinID uint) ([]*wallet.Wallet, error) {
	var _wallet []*wallet.Wallet
	var _table *wallet.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("coin_id = ?", coinID).Find(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletService) ScanWalletByUserID(tx *gorm.DB, userID uint) ([]*wallet.Wallet, error) {
	var _wallet []*wallet.Wallet
	var _table *wallet.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("user_id = ?", userID).Find(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletService) InsertWallet(tx *gorm.DB, _wallet *wallet.Wallet) (*wallet.Wallet, error) {
	var err error

	if err = tx.Table(_wallet.TableName()).Create(_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletService) UpdateWallet(tx *gorm.DB, id uint, address string) (*wallet.Wallet, error) {
	var _table *wallet.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("id = ?", id).Update("address", address).Error; err != nil {
		return nil, err
	}

	return w.FindWalletByID(tx, id)
}