package order

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderMatchingEvent struct {
	ID              uint            `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'"`
	OrderSymbolID   uint            ` gorm:"comment:'심볼 id'"`
	OrderIntervalID uint            `gorm:"comment:'시간대 id'"`
	Quantity        decimal.Decimal `gorm:"type:decimal(30,4);not null;comment:'수량'" sql:"type:decimal(30,4)"`
	UnitPrice       decimal.Decimal `gorm:"type:decimal(30,4);not null;comment:'단위 가격'" sql:"type:decimal(30,4)"`
	OrderType       string          `gorm:"not null;comment:'bid/ask';type:enum('BID', 'ASK')"`
	CreatedAt       time.Time       `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt       time.Time       `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (l *OrderMatchingEvent) Alias() string {
	return "order_matching_event ome"
}

func (l *OrderMatchingEvent) TableName() string {
	return "order_matching_event"
}
