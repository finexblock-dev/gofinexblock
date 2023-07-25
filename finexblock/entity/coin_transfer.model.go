package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type TransferType string

const (
	Deposit    TransferType = "DEPOSIT"
	Withdrawal TransferType = "WITHDRAWAL"
	Trade      TransferType = "TRADE"
)

type CoinTransfer struct {
	ID           uint            `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id"`
	WalletID     uint            `gorm:"comment:'지갑 id'" json:"walletId"`
	Amount       decimal.Decimal `gorm:"not null;type:decimal(30,4);comment:'수량'" sql:"type:decimal(30,4)" json:"amount"`
	TransferType TransferType    `gorm:"type:enum('DEPOSIT', 'WITHDRAWAL', 'TRADE');not null;comment:'입금/출금'" json:"transferType"`
	CreatedAt    time.Time       `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt    time.Time       `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`

	CoinTransaction   []*CoinTransaction   `gorm:"foreignKey:CoinTransferID;constraint:OnUpdate:CASCADE" json:"coinTransaction"`
	WithdrawalRequest []*WithdrawalRequest `gorm:"foreignKey:CoinTransferID;constraint:OnUpdate:CASCADE" json:"withdrawalRequest"`
	Wallet            *Wallet              `gorm:"foreignKey:WalletID;references:ID" json:"wallet"`
}

func (t *CoinTransfer) Alias() string {
	return "coin_transfer ct"
}

func (t *CoinTransfer) TableName() string {
	return "coin_transfer"
}