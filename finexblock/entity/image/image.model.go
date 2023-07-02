package image

import (
	"time"
)

type Image struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey;autoIncrement;not null;comment:'기본키'"`
	Key       string    `json:"key,omitempty" gorm:"not null;type:longtext;comment:'image key'"`
	Url       string    `json:"url,omitempty;" gorm:"not null;type:longtext;comment:'이미지 url'"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;comment:'생성일자';default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;comment:'수정일자';default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (i *Image) TableName() string {
	return "image"
}
