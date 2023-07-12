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
	ID           uint            `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'"`
	WalletID     uint            `gorm:"comment:'지갑 id'"`
	Amount       decimal.Decimal `gorm:"not null;type:decimal(30,4);comment:'수량'" sql:"type:decimal(30,4)"`
	TransferType TransferType    `gorm:"type:enum('DEPOSIT', 'WITHDRAWAL', 'TRADE');not null;comment:'입금/출금'"`
	CreatedAt    time.Time       `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt    time.Time       `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp"`

	CoinTransaction   []*CoinTransaction   `gorm:"foreignKey:CoinTransferID;constraint:OnUpdate:CASCADE"`
	WithdrawalRequest []*WithdrawalRequest `gorm:"foreignKey:CoinTransferID;constraint:OnUpdate:CASCADE"`
	Wallet            *Wallet              `gorm:"foreignKey:WalletID;references:ID"`
}

func (t *CoinTransfer) Alias() string {
	return "coin_transfer ct"
}

func (t *CoinTransfer) TableName() string {
	return "coin_transfer"
}
