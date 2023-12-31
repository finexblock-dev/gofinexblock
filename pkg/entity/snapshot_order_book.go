package entity

import (
	"time"
)

type SnapshotOrderBook struct {
	ID            uint      `gorm:"primary_key;auto_increment;comment:'기본키'" json:"id"`
	OrderSymbolID uint      `gorm:"not null;comment:'코인쌍';index" json:"orderSymbolId"` // Index added for foreign key
	BidOrderList  string    `gorm:"type:longtext;not null;comment:'매수주문리스트 문자열'" json:"bidOrderList"`
	AskOrderList  string    `gorm:"type:longtext;not null;comment:'매도주문리스트 문자열'" json:"askOrderList"`
	CreatedAt     time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'생성일자'" json:"createdAt"`
}

func (s *SnapshotOrderBook) Alias() string {
	return "snapshot_order_book sob"
}

func (s *SnapshotOrderBook) TableName() string {
	return "snapshot_order_book"
}