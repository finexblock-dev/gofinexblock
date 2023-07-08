package user

import "time"

type UserProfile struct {
	ID              uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	UserID          uint      `gorm:"comment:'유저 id'"`
	GradeID         uint      `gorm:"comment:'유저 등급 id';not null"`
	Nickname        string    `gorm:"comment:'유저 닉네임';not null"`
	Fullname        string    `gorm:"comment:'전체 이름, NICE 인증시 사용하는 컬럼';size:100"`
	Firstname       string    `gorm:"comment:'이름';size:100"`
	Lastname        string    `gorm:"comment:'성';size:100"`
	Birth           string    `gorm:"comment:'생년월일(ex. 20230101)';size:100"`
	Gender          string    `gorm:"comment:'성별';size:100"`
	ProfileImage    string    `gorm:"comment:'프로필 이미지 url';type:longtext"`
	ProfileImageKey string    `gorm:"comment:'프로필 이미지 url 키';type:longtext"`
	CountryCode     string    `gorm:"size:4;comment:'전화번호 앞자리'"`
	PhoneNumber     string    `gorm:"comment:'전화번호';type:longtext"`
	CreatedAt       time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt       time.Time `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt       time.Time `gorm:"comment:'삭제일자';index"`
}

func (p *UserProfile) Alias() string {
	return "user_profile up"
}

func (p *UserProfile) TableName() string {
	return "user_profile"
}
