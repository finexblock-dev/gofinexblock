package entity

import (
	"github.com/finexblock-dev/gofinexblock/pkg/types"
	"github.com/shopspring/decimal"
	"time"
)

type OrderBook struct {
	ID            uint              `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'" json:"id"`
	OrderSymbolID uint              `gorm:"not null;comment:'심볼 id'" json:"orderSymbolId"`
	UserID        uint              `gorm:"comment:'유저 id'" json:"userId"`
	OrderUUID     string            `gorm:"not null;unique;size:32;comment:'주문 uuid'" json:"orderUUID"`
	UnitPrice     decimal.Decimal   `gorm:"type:decimal(30,4);not null;comment:'단위가격'" sql:"type:decimal(30,4)" json:"unitPrice"`
	Quantity      decimal.Decimal   `gorm:"type:decimal(30,4);not null;comment:'수량'" sql:"type:decimal(30,4)" json:"quantity"`
	OrderType     types.OrderType   `gorm:"not null;type:enum('BID', 'ASK');comment:'bid or ask'" json:"orderType"`
	Status        types.OrderStatus `gorm:"not null;comment:'주문 상태';type:enum('CANCELLED', 'PLACED', 'FULFILLED', 'PARTIAL_FILLED');" json:"status"`
	CreatedAt     time.Time         `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt     time.Time         `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`

	OrderBookDifference []*OrderBookDifference `gorm:"foreignKey:OrderBookID;constraint:OnUpdate:CASCADE;" json:"orderBookDifference"`
}

func (b *OrderBook) Alias() string {
	return "order_book ob"
}

func (b *OrderBook) TableName() string {
	return "order_book"
}