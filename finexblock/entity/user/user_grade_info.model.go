package user

import (
	"github.com/shopspring/decimal"
	"time"
)

type UserGradeInfo struct {
	ID            uint            `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	Grade         string          `json:"grade,omitempty" gorm:"comment:'유저 등급';not null;type:enum('REGULAR','VIP1','VIP2','VIP3','VIP4','VIP5','VIP6','VIP7','VIP8','VIP9','VIP10')"`
	TakerTradeFee decimal.Decimal `json:"taker_trade_fee" gorm:"comment:'시장가 수수료';not null;type:decimal(6,5)"`
	MakerTradeFee decimal.Decimal `json:"maker_trade_fee" gorm:"comment:'지정가 수수료';not null;type:decimal(6,5)"`
	WithdrawFee   decimal.Decimal `json:"withdraw_fee" gorm:"comment:'출금 수수료';not null;type:decimal(6,5)"`
	WithdrawLimit decimal.Decimal `json:"withdraw_limit" gorm:"comment:'출금 한도';not null;type:decimal(30,4)"`
	CreatedAt     time.Time       `json:"created_at" gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time       `json:"updated_at" gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP"`
	DeletedAt     time.Time       `json:"deleted_at" gorm:"type:timestamp;comment:'삭제일자'"`
}

func (g *UserGradeInfo) Alias() string {
	return "user_grade_info ugi"
}

func (g *UserGradeInfo) TableName() string {
	return "user_grade_info"
}
