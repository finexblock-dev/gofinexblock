package entity

import "time"

type UserReferralCode struct {
	ID           uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	ID2          uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	UserID       uint      `gorm:"comment:'유저 id'"`
	ReferralCode uint      `gorm:"comment:'추천인 코드''"`
	CreatedAt    time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (r *UserReferralCode) Alias() string {
	return "user_referral_code urc"
}

func (r *UserReferralCode) TableName() string {
	return "user_referral_code"
}
