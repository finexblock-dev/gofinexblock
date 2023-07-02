package wallet

type CoinTransaction struct {
	ID             uint   `json:"id,omitempty" gorm:"primaryKey;autoIncrement;not null;comment:'기본키'"`
	CoinTransferID uint   `json:"coin_transfer_id,omitempty" gorm:"comment:'코인 전송 id'"`
	TxHash         string `json:"tx_hash,omitempty" gorm:"not null;comment:'트랜잭션 id'"`
}

func (c *CoinTransaction) Alias() string {
	return "coin_transaction ct"
}

func (c *CoinTransaction) TableName() string {
	return "coin_transaction"
}
