package user

import "time"

type UserVerification struct {
	ID            uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	UserID        uint      `gorm:"comment:'유저 id'"`
	PhoneVerified bool      `gorm:"comment: '전화번호 인증 여부'"`
	EmailVerified bool      `gorm:"comment: '이메일 인증 여부'"`
	OtpVerified   bool      `gorm:"comment: 'OTP 인증 여부'"`
	AdultVerified bool      `gorm:"comment: '성인 인증 여부'"`
	CreatedAt     time.Time `gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP"`
	DeletedAt     time.Time `gorm:"type:timestamp;comment:'삭제일자'"`
}

func (v *UserVerification) Alias() string {
	return "user_verification uv"
}

func (v *UserVerification) TableName() string {
	return "user_verification"
}
