package entity

type OrderSymbol struct {
	ID   uint   `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'"`
	Name string `gorm:"not null;comment:'심볼 이름(e.g. BTCETH)';size:32;unique;"`

	OrderMatchingEvent   []OrderMatchingEvent   `gorm:"foreignKey:OrderSymbolID;constraint:OnUpdate:CASCADE;"`
	OrderMatchingHistory []OrderMatchingHistory `gorm:"foreignKey:OrderSymbolID;constraint:OnUpdate:CASCADE;"`
	Chart                []Chart                `gorm:"foreignKey:OrderSymbolID;constraint:OnUpdate:CASCADE;"`
	OrderBook            []OrderBook            `gorm:"foreignKey:OrderSymbolID;constraint:OnUpdate:CASCADE;"`
}

func (s *OrderSymbol) Alias() string {
	return "order_symbol os"
}

func (s *OrderSymbol) TableName() string {
	return "order_symbol"
}
