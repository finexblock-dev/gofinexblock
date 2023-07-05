package wallet

type CoinTransaction struct {
	ID             uint   `gorm:"primaryKey;autoIncrement;not null;comment:'기본키'"`
	CoinTransferID uint   `gorm:"comment:'코인 전송 id'"`
	TxHash         string `gorm:"not null;comment:'트랜잭션 id'"`
}

func (c *CoinTransaction) Alias() string {
	return "coin_transaction ct"
}

func (c *CoinTransaction) TableName() string {
	return "coin_transaction"
}