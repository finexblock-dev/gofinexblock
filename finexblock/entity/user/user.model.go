package user

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/order"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"time"
)

type User struct {
	ID   uint   `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	UUID string `gorm:"not null;size:255;comment:'유저 uuid'"`

	IsEmailUser bool      `gorm:"not null;default:false;comment:'자체 회원가입 여부'"`
	UserType    string    `gorm:"not null;default:'NORMAL';comment:'유저 타입'"`
	IsBlock     bool      `gorm:"not null;default:false;comment:'블락 여부'"`
	CreatedAt   time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt   time.Time `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt   time.Time `gorm:"comment:'삭제일자';index"`

	Wallet               []wallet.Wallet   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	OrderBook            []order.OrderBook `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	OrderMatchingHistory []order.OrderBook `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`

	UserDormant          []UserDormant        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	UserDormantLog       []UserDormantLog     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	UserProfile          UserProfile          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	UserMemo             UserMemo             `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	UserEmailSignUp      UserEmailSignUp      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	UserSingleSignOnInfo UserSingleSignOnInfo `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Alias() string {
	return "user u"
}