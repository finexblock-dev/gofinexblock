package entity

type SmartContract struct {
	ID      uint64 `gorm:"primary_key;auto_increment;comment:'기본키'" json:"id"`
	CoinID  uint64 `gorm:"not null;comment:'코인 id'" json:"coinId"`
	Address string `gorm:"comment:'컨트랙트 주소';not null;type:longtext" json:"address"`
}

func (s *SmartContract) TableName() string {
	return "smart_contract"
}

func (s *SmartContract) Alias() string {
	return "smart_contract s"
}