package entity

import (
	"time"
)

type Coin struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id"`
	BlockchainID uint      `gorm:"not null;comment:'블록체인 ID'" json:"blockchainId"`
	Name         string    `gorm:"comment:'코인이름 (e.g. ETH)'" json:"name"`
	CreatedAt    time.Time `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`

	Wallet []*Wallet `gorm:"foreignKey:CoinID;constraint:OnUpdate:CASCADE;" json:"wallet"`
}

func (c *Coin) Alias() string {
	return "coin c"
}

func (c *Coin) TableName() string {
	return "coin"
}