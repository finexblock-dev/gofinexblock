package entity

type OrderSymbol struct {
	ID   uint   `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'" json:"id"`
	Name string `gorm:"not null;comment:'심볼 이름(e.g. BTCETH)';size:32;unique;" json:"name"`

	OrderMatchingEvent   []OrderMatchingEvent   `gorm:"foreignKey:OrderSymbolID;constraint:OnUpdate:CASCADE;" json:"orderMatchingEvent"`
	OrderMatchingHistory []OrderMatchingHistory `gorm:"foreignKey:OrderSymbolID;constraint:OnUpdate:CASCADE;" json:"orderMatchingHistory"`
	Chart                []Chart                `gorm:"foreignKey:OrderSymbolID;constraint:OnUpdate:CASCADE;" json:"chart"`
	OrderBook            []OrderBook            `gorm:"foreignKey:OrderSymbolID;constraint:OnUpdate:CASCADE;" json:"orderBook"`
}

func (s *OrderSymbol) Alias() string {
	return "order_symbol os"
}

func (s *OrderSymbol) TableName() string {
	return "order_symbol"
}