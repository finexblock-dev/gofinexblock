package user

import "time"

type UserDormant struct {
	ID              uint        `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	UserID          uint        `gorm:"comment:'유저 id'"`
	MetadataProfile interface{} `gorm:"type:json;comment:'프로필 메타데이터';not null;"`
	MetadataUser    interface{} `gorm:"type:json;comment:'유저 메타데이터';not null;"`
	CreatedAt       time.Time   `gorm:"type:timestamp;comment:' 생성일자';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time   `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP"`
	DeletedAt       time.Time   `gorm:"type:timestamp;comment:'삭제일자'"`
}

func (d *UserDormant) Alias() string {
	return "user_dormant ud"
}

func (d *UserDormant) TableName() string {
	return "user_dormant"
}
