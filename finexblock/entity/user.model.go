package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	ID   uint   `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UUID string `gorm:"not null;size:255;comment:'유저 uuid'" json:"uuid"`

	IsEmailUser bool      `gorm:"not null;default:false;comment:'자체 회원가입 여부'" json:"isEmailUser"`
	UserType    string    `gorm:"not null;default:'NORMAL';comment:'유저 타입'" json:"userType"`
	IsBlock     bool      `gorm:"not null;default:false;comment:'블락 여부'" json:"isBlock"`
	CreatedAt   time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`
	DeletedAt   time.Time `gorm:"comment:'삭제일자';index" json:"deletedAt"`

	Wallet               []Wallet    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;" json:"wallet"`
	OrderBook            []OrderBook `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;" json:"orderBook"`
	OrderMatchingHistory []OrderBook `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;" json:"orderMatchingHistory"`

	UserDormant          []UserDormant        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;" json:"userDormant"`
	UserDormantLog       []UserDormantLog     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;" json:"userDormantLog"`
	UserProfile          UserProfile          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;" json:"userProfile"`
	UserMemo             UserMemo             `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;" json:"userMemo"`
	UserEmailSignUp      UserEmailSignUp      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;" json:"userEmailSignUp"`
	UserSingleSignOnInfo UserSingleSignOnInfo `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;" json:"userSingleSignOnInfo"`
}

type UserMetadata struct {
	ID                uint            `json:"id" query:"id"`
	UUID              string          `json:"uuid" query:"uuid"`
	UserType          string          `json:"userType" query:"user_type"`
	Nickname          string          `json:"nickname" query:"nickname"`
	Fullname          string          `json:"fullname" query:"fullname"`
	PhoneNumber       string          `json:"phoneNumber" query:"phone_number"`
	BTC               decimal.Decimal `json:"btc" query:"btc"`
	IsBlock           bool            `json:"isBlock" query:"is_block"`
	IsDormant         bool            `json:"isDormant" query:"is_dormant"`
	IsMetaverseUser   bool            `json:"isMetaverseUser" query:"is_metaverse_user"`
	IsGoogleUser      bool            `json:"isGoogleUser" query:"is_google_user"`
	IsAppleUser       bool            `json:"isAppleUser" query:"is_apple_user"`
	IsEmailSignUpUser bool            `json:"isEmailSignUpUser" query:"is_email_sign_up_user"`
	CreatedAt         time.Time       `json:"createdAt" query:"created_at"`
	UpdatedAt         time.Time       `json:"updatedAt" query:"updated_at"`
	UserMemo          *UserMemo       `json:"userMemo" query:"user_memo"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Alias() string {
	return "user u"
}