package entity

import "time"

type UserPhoneCode struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	UserID    uint      `gorm:"comment:'유저 id'"`
	AuthCode  string    `gorm:"not null;comment:'코드';"`
	ExpiredAt time.Time `gorm:"type:timestamp;comment:'만료일자';not null"`
	CreatedAt time.Time `gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `gorm:"type:timestamp;comment:'삭제일자'"`
}

func (p *UserPhoneCode) Alias() string {
	return "user_phone_code upc"
}

func (p *UserPhoneCode) TableName() string {
	return "user_phone_code"
}
