package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type UserGradeInfo struct {
	ID            uint            `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	Grade         string          `gorm:"comment:'유저 등급';not null;type:enum('REGULAR','VIP1','VIP2','VIP3','VIP4','VIP5','VIP6','VIP7','VIP8','VIP9','VIP10')" json:"grade"`
	TakerTradeFee decimal.Decimal `gorm:"comment:'시장가 수수료';not null;type:decimal(6,5)" json:"takerTradeFee"`
	MakerTradeFee decimal.Decimal `gorm:"comment:'지정가 수수료';not null;type:decimal(6,5)" json:"makerTradeFee"`
	WithdrawFee   decimal.Decimal `gorm:"comment:'출금 수수료';not null;type:decimal(6,5)" json:"withdrawFee"`
	WithdrawLimit decimal.Decimal `gorm:"comment:'출금 한도';not null;type:decimal(30,4)" json:"withdrawLimit"`
	CreatedAt     time.Time       `gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt     time.Time       `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt     time.Time       `gorm:"type:timestamp;comment:'삭제일자'" json:"deletedAt"`
}

func (g *UserGradeInfo) Alias() string {
	return "user_grade_info ugi"
}

func (g *UserGradeInfo) TableName() string {
	return "user_grade_info"
}