package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Chart struct {
	ID              uint            `json:"id" gorm:"primaryKey;autoIncrement;not null;comment:'기본키'"`
	OrderSymbolID   uint            `json:"orderSymbolId" gorm:"not null;comment:'심볼 id'"`
	OrderIntervalID uint            `json:"orderIntervalId" gorm:"not null;comment:'인터벌 id'"`
	OpenPrice       decimal.Decimal `json:"openPrice" gorm:"type:decimal(30,4);not null;comment:'시작가'" sql:"type:decimal(30,4)"`
	LowPrice        decimal.Decimal `json:"lowPrice" gorm:"type:decimal(30,4);not null;comment:'최저가'" sql:"type:decimal(30,4)"`
	HighPrice       decimal.Decimal `json:"highPrice" gorm:"type:decimal(30,4);not null;comment:'최고가'" sql:"type:decimal(30,4)"`
	ClosePrice      decimal.Decimal `json:"closePrice" gorm:"type:decimal(30,4);not null;comment:'종가'" sql:"type:decimal(30,4)"`
	Volume          decimal.Decimal `json:"volume" gorm:"not null;comment:'거래량'" sql:"type:decimal(30,4)"`
	TradingValue    decimal.Decimal `json:"tradingValue" gorm:"not null;comment:'거래대금'" sql:"type:decimal(30,4)"`
	CreatedAt       time.Time       `json:"createdAt" gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt       time.Time       `json:"updatedAt" gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (c *Chart) Alias() string {
	return "chart c"
}

func (c *Chart) TableName() string {
	return "chart"
}