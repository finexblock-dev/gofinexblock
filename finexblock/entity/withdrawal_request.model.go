package entity

import (
	"errors"
	"github.com/shopspring/decimal"
	"time"
)

type WithdrawalStatus string

func (w WithdrawalStatus) String() string {
	return string(w)
}

func (w WithdrawalStatus) Validate() error {
	switch w {
	case SUBMITTED, APPROVED, CANCELED, REJECTED, PENDING, COMPLETED, FAILED:
		return nil
	}
	return errors.New("invalid withdrawal status")
}

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
	ID             uint             `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id"`
	CoinTransferID uint             `gorm:"comment:'선차감 id'" json:"coinTransferId"`
	ToAddress      string           `gorm:"comment:'출금 주소';not null;type:longtext" json:"toAddress"`
	Amount         decimal.Decimal  `gorm:"type:decimal(30,4);not null;comment:'출금량';" json:"amount"`
	Fee            decimal.Decimal  `gorm:"type:decimal(30,4);not null;comment:'수수료';" json:"fee"`
	Status         WithdrawalStatus `gorm:"not null;type:enum('SUBMITTED', 'APPROVED', 'CANCELED', 'REJECTED', 'PENDING', 'COMPLETED', 'FAILED');comment:'상태'" json:"status"`
	CreatedAt      time.Time        `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt      time.Time        `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`
	CoinTransfer   *CoinTransfer    `gorm:"foreignKey:CoinTransferID;references:ID" json:"coinTransfer"`
}

func (w *WithdrawalRequest) Alias() string {
	return "withdrawal_request wr"
}

func (w *WithdrawalRequest) TableName() string {
	return "withdrawal_request"
}