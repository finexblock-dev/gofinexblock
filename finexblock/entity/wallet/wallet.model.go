package wallet

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/user"
	"time"
)

type Wallet struct {
	ID        uint      `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'"`
	UserID    uint      `gorm:"comment:'유저 id'"`
	CoinID    uint      `gorm:"comment:'코인 id'"`
	Address   string    `gorm:"comment:'입금 주소';not null;type:longtext;"`
	CreatedAt time.Time `gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp"`

	User         *user.User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
	Coin         *Coin           `gorm:"foreignKey:CoinID;constraint:OnUpdate:CASCADE"`
	CoinTransfer []*CoinTransfer `gorm:"foreignKey:WalletID;constraint:OnUpdate:CASCADE"`
}

func (w *Wallet) Alias() string {
	return "wallet w"
}

func (w *Wallet) TableName() string {
	return "wallet"
}