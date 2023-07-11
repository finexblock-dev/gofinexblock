package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type WithdrawalStatus string

const (
	SUBMITTED WithdrawalStatus = "SUBMITTED"
	APPROVED  WithdrawalStatus = "APPROVED"
	CANCELED  WithdrawalStatus = "CANCELED"
	REJECTED  WithdrawalStatus = "REJECTED"
	PENDING   WithdrawalStatus = "PENDING"
	COMPLETED WithdrawalStatus = "COMPLETED"
	FAILED    WithdrawalStatus = "FAILED"
)

type WithdrawalRequest struct {
	ID             uint             `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'"`
	CoinTransferID uint             `gorm:"comment:'선차감 id'"`
	ToAddress      string           `gorm:"comment:'출금 주소';not null;type:longtext"`
	Amount         decimal.Decimal  `gorm:"type:decimal(30,4);not null;comment:'출금량';"`
	Fee            decimal.Decimal  `gorm:"type:decimal(30,4);not null;comment:'수수료';"`
	Status         WithdrawalStatus `gorm:"not null;type:enum('SUBMITTED', 'APPROVED', 'CANCELED', 'REJECTED', 'PENDING', 'COMPLETED', 'FAILED');comment:'상태'"`
	CreatedAt      time.Time        `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt      time.Time        `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp"`
	CoinTransfer   *CoinTransfer    `gorm:"foreignKey:CoinTransferID;references:ID"`
}

func (w *WithdrawalRequest) Alias() string {
	return "withdrawal_request wr"
}

func (w *WithdrawalRequest) TableName() string {
	return "withdrawal_request"
}
