package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) FindSmartContractByCoinID(tx *gorm.DB, coinID uint) (*wallet.SmartContract, error) {
	var smartContract *wallet.SmartContract
	var err error

	if err = tx.Table(smartContract.TableName()).Where("coin_id = ?", coinID).First(&smartContract).Error; err != nil {
		return nil, err
	}

	return smartContract, nil
}

func (w *walletService) FindSmartContractByID(tx *gorm.DB, id uint) (*wallet.SmartContract, error) {
	var smartContract *wallet.SmartContract
	var err error

	if err = tx.Table(smartContract.TableName()).Where("id = ?", id).First(&smartContract).Error; err != nil {
		return nil, err
	}

	return smartContract, nil
}
