package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) FindWithdrawalRequestByID(tx *gorm.DB, id uint) (*wallet.WithdrawalRequest, error) {
	var _withdrawalRequest *wallet.WithdrawalRequest
	var err error

	if err = tx.Table(_withdrawalRequest.TableName()).Where("id = ?", id).First(_withdrawalRequest).Error; err != nil {
		return nil, err
	}

	return _withdrawalRequest, nil
}

func (w *walletService) ScanWithdrawalRequestByStatus(tx *gorm.DB, cond wallet.WithdrawalStatus) ([]*wallet.WithdrawalRequest, error) {
	var _withdrawalRequest []*wallet.WithdrawalRequest
	var _table *wallet.WithdrawalRequest
	var err error

	if err = tx.Table(_table.TableName()).Where("status = ?", cond).Find(_withdrawalRequest).Error; err != nil {
		return nil, err
	}

	return _withdrawalRequest, nil
}

func (w *walletService) UpdateWithdrawalRequest(tx *gorm.DB, id uint, state wallet.WithdrawalStatus) (*wallet.WithdrawalRequest, error) {
	var _withdrawalRequest *wallet.WithdrawalRequest
	var err error

	if err = tx.Table(_withdrawalRequest.TableName()).Where("id = ?", id).Update("status", state).Error; err != nil {
		return nil, err
	}

	return w.FindWithdrawalRequestByID(tx, id)
}