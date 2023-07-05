package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) FindBlockchainByName(tx *gorm.DB, name string) (*wallet.Blockchain, error) {
	var blockchain *wallet.Blockchain
	var err error

	if err = tx.Table(blockchain.TableName()).Where("name = ?", name).First(&blockchain).Error; err != nil {
		return nil, err
	}

	return blockchain, nil
}

func (w *walletService) FindBlockchainByID(tx *gorm.DB, id uint) (*wallet.Blockchain, error) {
	var blockchain *wallet.Blockchain
	var err error

	if err = tx.Table(blockchain.TableName()).Where("id = ?", id).First(&blockchain).Error; err != nil {
		return nil, err
	}

	return blockchain, nil
}