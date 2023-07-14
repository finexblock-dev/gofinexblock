package entity

import (
	"github.com/shopspring/decimal"
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

	Wallet               []Wallet    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	OrderBook            []OrderBook `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	OrderMatchingHistory []OrderBook `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`

	UserDormant          []UserDormant        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	UserDormantLog       []UserDormantLog     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	UserProfile          UserProfile          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	UserMemo             UserMemo             `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	UserEmailSignUp      UserEmailSignUp      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
	UserSingleSignOnInfo UserSingleSignOnInfo `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;"`
}

type UserMetadata struct {
	ID                uint            `json:"id" query:"id"`
	UUID              string          `json:"uuid" query:"uuid"`
	UserType          string          `json:"user_type" query:"user_type"`
	Nickname          string          `json:"nickname" query:"nickname"`
	Fullname          string          `json:"fullname" query:"fullname"`
	PhoneNumber       string          `json:"phone_number" query:"phone_number"`
	BTC               decimal.Decimal `json:"btc" query:"btc"`
	IsBlock           bool            `json:"is_block" query:"is_block"`
	IsDormant         bool            `json:"is_dormant" query:"is_dormant"`
	IsMetaverseUser   bool            `json:"is_metaverse_user" query:"is_metaverse_user"`
	IsGoogleUser      bool            `json:"is_google_user" query:"is_google_user"`
	IsAppleUser       bool            `json:"is_apple_user" query:"is_apple_user"`
	IsEmailSignUpUser bool            `json:"is_email_sign_up_user" query:"is_email_sign_up_user"`
	CreatedAt         time.Time       `json:"created_at" query:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at" query:"updated_at"`
	UserMemo          *UserMemo       `json:"user_memo" query:"user_memo"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Alias() string {
	return "user u"
}