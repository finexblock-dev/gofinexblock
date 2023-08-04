package entity

import (
	"time"
)

type Announcement struct {
	ID           uint                  `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	CategoryID   uint                  `gorm:"not null;comment:'공지 유형'" json:"categoryId"`
	KoreanTitle  string                `json:"koreanTitle" gorm:"type:longtext;not null;comment:'제목';"`
	EnglishTitle string                `json:"englishTitle" gorm:"type:longtext;not null;comment:'제목';"`
	ChineseTitle string                `json:"chineseTitle" gorm:"type:longtext;not null;comment:'제목';"`
	Korean       string                `json:"korean" gorm:"type:longtext;comment:'한글판';column:kor"`
	English      string                `json:"english" gorm:"type:longtext;comment:'영어판';column:eng"`
	Chinese      string                `json:"chinese" gorm:"type:longtext;comment:'중국판';column:cn"`
	Visible      bool                  `json:"visible" gorm:"comment:'가시 여부';not null;default:true"`
	Pinned       bool                  `json:"pinned" gorm:"comment:'고정 여부';not null;default:false"`
	CreatedAt    time.Time             `json:"createdAt" gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;"`
	UpdatedAt    time.Time             `json:"updatedAt" gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;"`
	Category     *AnnouncementCategory `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
}

func (a *Announcement) TableName() string {
	return "announcement"
}