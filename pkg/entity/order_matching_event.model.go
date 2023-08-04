package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderMatchingEvent struct {
	ID              uint            `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'" json:"id"`
	OrderSymbolID   uint            ` gorm:"comment:'심볼 id'" json:"orderSymbolId"`
	OrderIntervalID uint            `gorm:"comment:'시간대 id'" json:"orderIntervalId"`
	Quantity        decimal.Decimal `gorm:"type:decimal(30,4);not null;comment:'수량'" sql:"type:decimal(30,4)" json:"quantity"`
	UnitPrice       decimal.Decimal `gorm:"type:decimal(30,4);not null;comment:'단위 가격'" sql:"type:decimal(30,4)" json:"unitPrice"`
	OrderType       string          `gorm:"not null;comment:'bid/ask';type:enum('BID', 'ASK')" json:"orderType"`
	CreatedAt       time.Time       `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt       time.Time       `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`
}

func (l *OrderMatchingEvent) Alias() string {
	return "order_matching_event ome"
}

func (l *OrderMatchingEvent) TableName() string {
	return "order_matching_event"
}