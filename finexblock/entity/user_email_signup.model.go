package entity

import "time"

type UserEmailSignUp struct {
	ID           uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID       uint      `gorm:"comment:'유저 id'" json:"userId"`
	Email        string    `gorm:"size:100;comment:'이메일';not null" json:"email"`
	Password     string    `gorm:"size:255;comment:'패스워드';not null;" json:"password"`
	PasswordSalt string    `gorm:"size:255;comment:'솔트';not null;" json:"passwordSalt"`
	CreatedAt    time.Time `gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt    time.Time `gorm:"type:timestamp;comment:'삭제일자'" json:"deletedAt"`
}

func (e *UserEmailSignUp) Alias() string {
	return "user_email_signup ues"
}

func (e *UserEmailSignUp) TableName() string {
	return "user_email_signup"
}