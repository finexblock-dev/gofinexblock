package entity

import (
	"time"
)

type Blockchain struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'"`
	Name      string    `gorm:"type:longtext;not null;comment:'체인 이름'"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:'생성일자';type:timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;comment:'수정일자';type:timestamp"`

	Coin []Coin `gorm:"foreignKey:BlockchainID;constraint:OnUpdate:CASCADE;"`
}

func (b *Blockchain) Alias() string {
	return "blockchain b"
}

func (b *Blockchain) TableName() string {
	return "blockchain"
}
