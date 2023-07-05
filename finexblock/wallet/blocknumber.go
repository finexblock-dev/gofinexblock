package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func (w *walletService) FindBlockNumberByCoinID(tx *gorm.DB, coinID uint) (*wallet.BlockNumber, error) {
	var blockNumber *wallet.BlockNumber
	var err error

	if err = tx.Table(blockNumber.TableName()).Where("coin_id = ?", coinID).First(&blockNumber).Error; err != nil {
		return nil, err
	}

	return blockNumber, nil
}

func (w *walletService) FindBlockNumberByID(tx *gorm.DB, id uint) (*wallet.BlockNumber, error) {
	var blockNumber *wallet.BlockNumber
	var err error

	if err = tx.Table(blockNumber.TableName()).Where("id = ?", id).First(&blockNumber).Error; err != nil {
		return nil, err
	}

	return blockNumber, nil
}

func (w *walletService) UpdateBlockNumber(tx *gorm.DB, coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (*wallet.BlockNumber, error) {
	var err error
	var table *wallet.BlockNumber

	if err = tx.Table(table.TableName()).Update("from_block", fromBlockNumber).Update("to_block", toBlockNumber).Error; err != nil {
		return nil, err
	}

	return table, nil
}