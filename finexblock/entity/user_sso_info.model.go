package entity

import "time"

type SSOType string

const (
	Metaverse SSOType = "METAVERSE"
	Apple     SSOType = "APPLE"
	Google    SSOType = "GOOGLE"
)

type UserSingleSignOnInfo struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	UserID    uint      `gorm:"comment:'유저 id'"`
	SSOType   SSOType   `gorm:"comment:'SSO 타입';type:enum('APPLE', 'GOOGLE', 'METAVERSE');not null"`
	AuthCode  string    `gorm:"not null;comment:'코드';"`
	Email     string    `gorm:"size:100;comment:'이메일';not null"`
	ExpiredAt time.Time `gorm:"type:timestamp;comment:'만료일자';not null"`
	CreatedAt time.Time `gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `gorm:"type:timestamp;comment:'삭제일자'"`
}

func (S *UserSingleSignOnInfo) Alias() string {
	return "user_sso_info usi"
}

func (S *UserSingleSignOnInfo) TableName() string {
	return "user_sso_info"
}
