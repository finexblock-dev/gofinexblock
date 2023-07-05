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

func (w *walletService) ScanWithdrawalRequestByStatus(tx *gorm.DB, status wallet.WithdrawalStatus) ([]*wallet.WithdrawalRequest, error) {
	var _withdrawalRequest []*wallet.WithdrawalRequest
	var table *wallet.WithdrawalRequest
	var err error

	if err = tx.Table(table.TableName()).Where("status = ?", status).Find(_withdrawalRequest).Error; err != nil {
		return nil, err
	}

	return _withdrawalRequest, nil
}

func (w *walletService) ScanWithdrawalRequestByCond(tx *gorm.DB, coinID uint, status wallet.WithdrawalStatus) ([]*wallet.WithdrawalRequest, error) {
	var withdrawalRequests []*wallet.WithdrawalRequest
	var table *wallet.WithdrawalRequest

	err := tx.Table(table.TableName()).
		Select("withdrawal_request.*").
		Joins("join coin_transfer on coin_transfer.id = withdrawal_request.coin_transfer_id").
		Joins("join wallet on wallet.id = coin_transfer.wallet_id").
		Where("wallet.coin_id = ?", coinID).
		Find(&withdrawalRequests).Error

	if err != nil {
		return nil, err
	}

	return withdrawalRequests, nil
}

func (w *walletService) UpdateWithdrawalRequest(tx *gorm.DB, id uint, state wallet.WithdrawalStatus) (*wallet.WithdrawalRequest, error) {
	var _withdrawalRequest *wallet.WithdrawalRequest
	var err error

	if err = tx.Table(_withdrawalRequest.TableName()).Where("id = ?", id).Update("status", state).Error; err != nil {
		return nil, err
	}

	return w.FindWithdrawalRequestByID(tx, id)
}