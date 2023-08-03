package entity

import (
	"errors"
	"time"
)

type TransactionStatus string

func (t TransactionStatus) String() string {
	return string(t)
}

func (t TransactionStatus) Validate() error {
	switch t {
	case INIT, DONE, REVERT:
		return nil
	}
	return errors.New("invalid transaction status")
}

const (
	INIT   TransactionStatus = "INIT"
	DONE   TransactionStatus = "DONE"
	REVERT TransactionStatus = "REVERT"
)

type CoinTransaction struct {
	ID             uint              `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id"`
	CoinTransferID uint              `gorm:"comment:'코인 전송 id'" json:"coinTransferId"`
	TxHash         string            `gorm:"not null;comment:'트랜잭션 id'" json:"txHash"`
	Status         TransactionStatus `gorm:"not null;type:enum('INIT', 'DONE', 'REVERT');default:'INIT';comment:'상태'" json:"status"`
	CreatedAt      time.Time         `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt      time.Time         `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`
}

func (c *CoinTransaction) Alias() string {
	return "coin_transaction ct"
}

func (c *CoinTransaction) TableName() string {
	return "coin_transaction"
}