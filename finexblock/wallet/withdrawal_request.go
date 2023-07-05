package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"gorm.io/gorm"
)

func (w *walletService) ScanWithdrawalRequestByStatus(cond wallet.WithdrawalStatus) ([]*wallet.WithdrawalRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (w *walletService) UpdateWithdrawalRequest(tx *gorm.DB, id uint, state wallet.WithdrawalStatus) error {
	//TODO implement me
	panic("implement me")
}