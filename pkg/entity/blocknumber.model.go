package entity

import "github.com/shopspring/decimal"

type BlockNumber struct {
	ID        uint64          `gorm:"primary_key;auto_increment;comment:'기본키'" json:"id"`
	CoinID    uint64          `gorm:"not null;comment:'코인 id'" json:"coinId"`
	FromBlock decimal.Decimal `gorm:"type:decimal(65,30);not null;comment:'검색 시작 블록'" json:"fromBlock"`
	ToBlock   decimal.Decimal `gorm:"type:decimal(65,30);not null;comment:'검색 끝나는 블록'" json:"toBlock"`
}

func (b *BlockNumber) TableName() string {
	return "blocknumber"
}

func (b *BlockNumber) Alias() string {
	return "blocknumber b"
}