package entity

type UserStarredTradingPair struct {
	ID            uint `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID        uint `gorm:"comment:'유저 id'" json:"userId"`
	OrderSymbolId uint `gorm:"comment: '거래쌍 id'" json:"orderSymbolId"`
}

func (s *UserStarredTradingPair) Alias() string {
	return "user_starred_trading_pair ustp"
}

func (s *UserStarredTradingPair) TableName() string {
	return "user_starred_trading_pair"
}