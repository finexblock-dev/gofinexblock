package user

import "time"

type UserFirebaseToken struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID    uint      `json:"user_id,omitempty" gorm:"comment:'유저 id'"`
	Token     string    `json:"token" gorm:"comment: 'firebase 토큰'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" gorm:"type:timestamp;comment:'삭제일자'"`
}

func (f *UserFirebaseToken) Alias() string {
	return "user_firebase_token uft"
}

func (f *UserFirebaseToken) TableName() string {
	return "user_firebase_token"
}