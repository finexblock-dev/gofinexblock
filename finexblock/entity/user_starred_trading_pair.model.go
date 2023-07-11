package entity

type UserStarredTraddingPair struct {
	ID            uint `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	UserID        uint `gorm:"comment:'유저 id'"`
	OrderSymbolId uint `gorm:"comment: '거래쌍 id'"`
}

func (s *UserStarredTraddingPair) Alias() string {
	return "user_starred_trading_pair ustp"
}

func (s *UserStarredTraddingPair) TableName() string {
	return "user_starred_trading_pair"
}
