package entity

import "time"

type UserDormantLog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID    uint      `gorm:"comment:'유저 id'" json:"userId"`
	Type      string    `gorm:"comment:'전환타입 - 휴면 전환, 휴면 해제, 개인정보제거';type:enum('DORMANT', 'EXIT_DORMANT', 'DELETED');not null" json:"type"`
	CreatedAt time.Time `gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt time.Time `gorm:"type:timestamp;comment:'삭제일자'" json:"deletedAt"`
}

func (d *UserDormantLog) Alias() string {
	return "user_dormant_log udl"
}

func (d *UserDormantLog) TableName() string {
	return "user_dormant_log"
}