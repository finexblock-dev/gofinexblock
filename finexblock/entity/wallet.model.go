package entity

import (
	"time"
)

type Wallet struct {
	ID        uint      `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'" json:"id"`
	UserID    uint      `gorm:"comment:'유저 id'" json:"userId"`
	CoinID    uint      `gorm:"comment:'코인 id'" json:"coinId"`
	Address   string    `gorm:"comment:'입금 주소';not null;type:longtext;" json:"address"`
	CreatedAt time.Time `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`

	CoinTransfer []*CoinTransfer `gorm:"foreignKey:WalletID;constraint:OnUpdate:CASCADE" json:"coinTransfer"`
}

func (w *Wallet) Alias() string {
	return "wallet w"
}

func (w *Wallet) TableName() string {
	return "wallet"
}