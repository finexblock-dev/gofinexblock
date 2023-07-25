package entity

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/shopspring/decimal"
	"time"
)

type OrderMatchingHistory struct {
	ID             uint            `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'" json:"id"`
	UserID         uint            `gorm:"not null;comment:'유저 id'" json:"userId"`
	OrderSymbolID  uint            `gorm:"not null;comment:'심볼 id'" json:"orderSymbolId"`
	OrderUUID      string          `gorm:"not null;comment:'주문 uuid';size:32" json:"orderUUID"`
	FilledQuantity decimal.Decimal `gorm:"not null;type:decimal(30,4);comment:'체결 수량'" sql:"type:decimal(30,4)" json:"filledQuantity"`
	UnitPrice      decimal.Decimal `gorm:"not null;type:decimal(30,4);comment:'단위 가격'" sql:"type:decimal(30,4)" json:"unitPrice"`
	Fee            decimal.Decimal `gorm:"not null;type:decimal(30,4);comment:'수수료'" sql:"type:decimal(30,4)" json:"fee"`
	OrderType      types.OrderType `gorm:"not null;comment:'bid/ask';type:enum('BID', 'ASK')" json:"orderType"`
	CreatedAt      time.Time       `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt      time.Time       `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`
}

func (m *OrderMatchingHistory) Alias() string {
	return "order_matching_history omh"
}

func (m *OrderMatchingHistory) TableName() string {
	return "order_matching_history"
}