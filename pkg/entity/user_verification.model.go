package entity

import "time"

type UserVerification struct {
	ID            uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID        uint      `gorm:"comment:'유저 id'" json:"userId"`
	PhoneVerified bool      `gorm:"comment: '전화번호 인증 여부'" json:"phoneVerified"`
	EmailVerified bool      `gorm:"comment: '이메일 인증 여부'" json:"emailVerified"`
	OtpVerified   bool      `gorm:"comment: 'OTP 인증 여부'" json:"otpVerified"`
	AdultVerified bool      `gorm:"comment: '성인 인증 여부'" json:"adultVerified"`
	CreatedAt     time.Time `gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt     time.Time `gorm:"type:timestamp;comment:'삭제일자'" json:"deletedAt"`
}

func (v *UserVerification) Alias() string {
	return "user_verification uv"
}

func (v *UserVerification) TableName() string {
	return "user_verification"
}