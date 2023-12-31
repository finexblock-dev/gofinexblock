package entity

type AnnouncementCategory struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	KoreanType  string `json:"koreanType" gorm:"type:longtext;not null;column:ko_type"`
	EnglishType string `json:"englishType" gorm:"type:longtext;not null;column:en_type"`
	ChineseType string `json:"chineseType" gorm:"type:longtext;not null;column:cn_type"`
}

func (c *AnnouncementCategory) TableName() string {
	return "announcement_category"
}