package order

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderBook struct {
	ID            uint            `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'"`
	OrderSymbolID uint            `gorm:"not null;comment:'심볼 id'"`
	UserID        uint            `gorm:"comment:'유저 id'"`
	OrderUUID     string          `json:"order_uuid" gorm:"not null;unique;size:32;comment:'주문 uuid'"`
	UnitPrice     decimal.Decimal `gorm:"type:decimal(30,4);not null;comment:'단위가격'" sql:"type:decimal(30,4)"`
	Quantity      decimal.Decimal `gorm:"type:decimal(30,4);not null;comment:'수량'" sql:"type:decimal(30,4)"`
	OrderType     string          `gorm:"not null;type:enum('BID', 'ASK');comment:'bid or ask'"`
	Status        string          `gorm:"not null;comment:'주문 상태';type:enum('CANCELLED', 'PLACED', 'FULFILLED', 'PARTIAL_FILLED');"`
	CreatedAt     time.Time       `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt     time.Time       `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp"`

	OrderBookDifference []*OrderBookDifference `gorm:"foreignKey:OrderBookID;constraint:OnUpdate:CASCADE;"`
}

func (b *OrderBook) Alias() string {
	return "order_book ob"
}

func (b *OrderBook) TableName() string {
	return "order_book"
}
