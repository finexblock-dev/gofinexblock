package entity

import "time"

type UserReferralCode struct {
	ID           uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID       uint      `gorm:"comment:'유저 id'" json:"userId"`
	ReferralCode uint      `gorm:"comment:'추천인 코드''" json:"referralCode"`
	CreatedAt    time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
}

func (r *UserReferralCode) Alias() string {
	return "user_referral_code urc"
}

func (r *UserReferralCode) TableName() string {
	return "user_referral_code"
}