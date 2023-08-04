package entity

import "time"

type UserProfile struct {
	ID              uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID          uint      `gorm:"comment:'유저 id'" json:"userId"`
	GradeID         uint      `gorm:"comment:'유저 등급 id';not null" json:"gradeId"`
	Nickname        string    `gorm:"comment:'유저 닉네임';not null" json:"nickname"`
	Fullname        string    `gorm:"comment:'전체 이름, NICE 인증시 사용하는 컬럼';size:100" json:"fullname"`
	Firstname       string    `gorm:"comment:'이름';size:100" json:"firstname"`
	Lastname        string    `gorm:"comment:'성';size:100" json:"lastname"`
	Birth           string    `gorm:"comment:'생년월일(ex. 20230101)';size:100" json:"birth"`
	Gender          string    `gorm:"comment:'성별';size:100" json:"gender"`
	ProfileImage    string    `gorm:"comment:'프로필 이미지 url';type:longtext" json:"profileImage"`
	ProfileImageKey string    `gorm:"comment:'프로필 이미지 url 키';type:longtext" json:"profileImageKey"`
	CountryCode     string    `gorm:"size:4;comment:'전화번호 앞자리'" json:"countryCode"`
	PhoneNumber     string    `gorm:"comment:'전화번호';type:longtext" json:"phoneNumber"`
	CreatedAt       time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`
	DeletedAt       time.Time `gorm:"comment:'삭제일자';index" json:"deletedAt"`
}

func (p *UserProfile) Alias() string {
	return "user_profile up"
}

func (p *UserProfile) TableName() string {
	return "user_profile"
}