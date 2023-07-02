package order

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderMatchingHistory struct {
	ID             uint            `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'"`
	UserID         uint            `gorm:"not null;comment:'유저 id'"`
	OrderSymbolID  uint            `gorm:"not null;comment:'심볼 id'"`
	OrderUUID      string          `gorm:"not null;comment:'주문 uuid';size:32"`
	FilledQuantity decimal.Decimal `gorm:"not null;type:decimal(30,4);comment:'체결 수량'" sql:"type:decimal(30,4)"`
	UnitPrice      decimal.Decimal `gorm:"not null;type:decimal(30,4);comment:'단위 가격'" sql:"type:decimal(30,4)"`
	Fee            decimal.Decimal `gorm:"not null;type:decimal(30,4);comment:'수수료'" sql:"type:decimal(30,4)"`
	OrderType      string          `gorm:"not null;comment:'bid/ask';type:enum('BID', 'ASK')"`
	CreatedAt      time.Time       `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt      time.Time       `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (m *OrderMatchingHistory) Alias() string {
	return "order_matching_history omh"
}

func (m *OrderMatchingHistory) TableName() string {
	return "order_matching_history"
}
