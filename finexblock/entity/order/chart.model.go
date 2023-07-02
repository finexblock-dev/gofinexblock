package order

import (
	"github.com/shopspring/decimal"
	"time"
)

type Chart struct {
	ID              uint            `json:"id,omitempty" gorm:"primaryKey;autoIncrement;not null;comment:'기본키'"`
	OrderSymbolID   uint            `json:"order_symbol_id,omitempty" gorm:"not null;comment:'심볼 id'"`
	OrderIntervalID uint            `json:"order_interval_id,omitempty" gorm:"not null;comment:'인터벌 id'"`
	OpenPrice       decimal.Decimal `json:"open_price,omitempty" gorm:"type:decimal(30,4);not null;comment:'시작가'" sql:"type:decimal(30,4)"`
	LowPrice        decimal.Decimal `json:"low_price,omitempty" gorm:"type:decimal(30,4);not null;comment:'최저가'" sql:"type:decimal(30,4)"`
	HighPrice       decimal.Decimal `json:"high_price,omitempty" gorm:"type:decimal(30,4);not null;comment:'최고가'" sql:"type:decimal(30,4)"`
	ClosePrice      decimal.Decimal `json:"close_price,omitempty" gorm:"type:decimal(30,4);not null;comment:'종가'" sql:"type:decimal(30,4)"`
	Volume          decimal.Decimal `json:"volume,omitempty" gorm:"not null;comment:'거래량'" sql:"type:decimal(30,4)"`
	TradingValue    decimal.Decimal `json:"trading_value,omitempty" gorm:"not null;comment:'거래대금'" sql:"type:decimal(30,4)"`
	CreatedAt       time.Time       `json:"created_at" gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (c *Chart) Alias() string {
	return "chart c"
}

func (c *Chart) TableName() string {
	return "chart"
}
