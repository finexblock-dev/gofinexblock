package entity

import "time"

type UserPhoneCode struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID    uint      `gorm:"comment:'유저 id'" json:"userId"`
	AuthCode  string    `gorm:"not null;comment:'코드';" json:"authCode"`
	ExpiredAt time.Time `gorm:"type:timestamp;comment:'만료일자';not null" json:"expiredAt"`
	CreatedAt time.Time `gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt time.Time `gorm:"type:timestamp;comment:'삭제일자'" json:"deletedAt"`
}

func (p *UserPhoneCode) Alias() string {
	return "user_phone_code upc"
}

func (p *UserPhoneCode) TableName() string {
	return "user_phone_code"
}