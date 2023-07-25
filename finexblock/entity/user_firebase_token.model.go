package entity

import "time"

type UserFirebaseToken struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID    uint      `gorm:"comment:'유저 id'" json:"userId"`
	Token     string    `gorm:"comment: 'firebase 토큰'" json:"token"`
	CreatedAt time.Time `gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt time.Time `gorm:"type:timestamp;comment:'삭제일자'" json:"deletedAt"`
}

func (f *UserFirebaseToken) Alias() string {
	return "user_firebase_token uft"
}

func (f *UserFirebaseToken) TableName() string {
	return "user_firebase_token"
}