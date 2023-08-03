package entity

import (
	"github.com/finexblock-dev/gofinexblock/pkg/types"
	"github.com/shopspring/decimal"
	"time"
)

type OrderBookDifference struct {
	ID          uint            `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'" json:"id"`
	OrderBookID uint            `gorm:"not null;comment:'호가 id'" json:"orderBookId"`
	Diff        decimal.Decimal `gorm:"type:decimal(30,4);not null;comment:'변동량';" sql:"type:decimal(30,4)" json:"diff"`
	Reason      types.Reason    `gorm:"type:enum('CANCEL', 'PLACE', 'FILL');not null;comment:'변동 사유';" json:"reason"`
	CreatedAt   time.Time       `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt   time.Time       `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`
}

func (p *OrderBookDifference) Alias() string {
	return "order_book_difference obd"
}

func (p *OrderBookDifference) TableName() string {
	return "order_book_difference"
}