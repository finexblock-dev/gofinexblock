package entity

import "time"

type UserDormant struct {
	ID              uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID          uint      `gorm:"comment:'유저 id'" json:"userId"`
	MetadataProfile Metadata  `gorm:"type:json;comment:'프로필 메타데이터';not null;" json:"metadataProfile"`
	MetadataUser    Metadata  `gorm:"type:json;comment:'유저 메타데이터';not null;" json:"metadataUser"`
	CreatedAt       time.Time `gorm:"type:timestamp;comment:' 생성일자';not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt       time.Time `gorm:"type:timestamp;comment:'삭제일자'" json:"deletedAt"`
}

func (d *UserDormant) Alias() string {
	return "user_dormant ud"
}

func (d *UserDormant) TableName() string {
	return "user_dormant"
}
