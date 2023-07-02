package user

import "time"

type UserEmailSignUp struct {
	ID           uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID       uint      `json:"user_id,omitempty" gorm:"comment:'유저 id'"`
	Email        string    `json:"email,omitempty" gorm:"size:100;comment:'이메일';not null"`
	Password     string    `json:"password,omitempty" gorm:"size:255;comment:'패스워드';not null;"`
	PasswordSalt string    `json:"password_salt,omitempty" gorm:"size:255;comment:'솔트';not null;"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP"`
	DeletedAt    time.Time `json:"deleted_at" gorm:"type:timestamp;comment:'삭제일자'"`
}

func (e *UserEmailSignUp) Alias() string {
	return "user_email_signup ues"
}

func (e *UserEmailSignUp) TableName() string {
	return "user_email_signup"
}
