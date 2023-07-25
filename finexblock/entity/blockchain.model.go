package entity

import (
	"time"
)

type Blockchain struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id"`
	Name      string    `gorm:"type:longtext;not null;comment:'체인 이름'" json:"name"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:'생성일자';type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:'수정일자';type:timestamp" json:"updatedAt"`

	Coin []Coin `gorm:"foreignKey:BlockchainID;constraint:OnUpdate:CASCADE;" json:"coin"`
}

func (b *Blockchain) Alias() string {
	return "blockchain b"
}

func (b *Blockchain) TableName() string {
	return "blockchain"
}