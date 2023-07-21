package entity

import (
	"time"
)

type Coin struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'"`
	BlockchainID uint      `gorm:"not null;comment:'블록체인 ID'"`
	Name         string    `gorm:"comment:'코인이름 (e.g. ETH)'"`
	CreatedAt    time.Time `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt    time.Time `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp"`

	Wallet []*Wallet `gorm:"foreignKey:CoinID;constraint:OnUpdate:CASCADE;"`
}

func (c *Coin) Alias() string {
	return "coin c"
}

func (c *Coin) TableName() string {
	return "coin"
}