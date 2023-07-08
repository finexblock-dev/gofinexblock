package user

import "time"

type UserEmailSignUp struct {
	ID           uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	UserID       uint      `gorm:"comment:'유저 id'"`
	Email        string    `gorm:"size:100;comment:'이메일';not null"`
	Password     string    `gorm:"size:255;comment:'패스워드';not null;"`
	PasswordSalt string    `gorm:"size:255;comment:'솔트';not null;"`
	CreatedAt    time.Time `gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP"`
	DeletedAt    time.Time `gorm:"type:timestamp;comment:'삭제일자'"`
}

func (e *UserEmailSignUp) Alias() string {
	return "user_email_signup ues"
}

func (e *UserEmailSignUp) TableName() string {
	return "user_email_signup"
}
