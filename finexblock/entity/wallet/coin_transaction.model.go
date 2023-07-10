package wallet

import "time"

type TransactionStatus string

const (
	INIT   TransactionStatus = "INIT"
	DONE   TransactionStatus = "DONE"
	REVERT TransactionStatus = "REVERT"
)

type CoinTransaction struct {
	ID             uint              `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'"`
	CoinTransferID uint              `gorm:"comment:'코인 전송 id'"`
	TxHash         string            `gorm:"not null;comment:'트랜잭션 id'"`
	Status         TransactionStatus `gorm:"not null;type:enum('INIT', 'DONE', 'REVERT');default:'INIT';comment:'상태'"`
	CreatedAt      time.Time         `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt      time.Time         `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (c *CoinTransaction) Alias() string {
	return "coin_transaction ct"
}

func (c *CoinTransaction) TableName() string {
	return "coin_transaction"
}