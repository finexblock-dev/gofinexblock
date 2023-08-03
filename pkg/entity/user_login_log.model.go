package entity

import "time"

type UserLoginLog struct {
	ID                uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID            uint      `gorm:"comment:'유저 id'" json:"userId"`
	IP                string    `gorm:"comment:'아이피';not null;size:200;" json:"ip"`
	Device            string    `gorm:"comment:'로그인 시도한 기기';not null;size:100" json:"device"`
	LoginMethod       string    `gorm:"comment:'로그인 방법';size:100;not null" json:"loginMethod"`
	IsSuccess         bool      `gorm:"comment:'로그인 성패 여부';type:tinyint(1);not null;default:false" json:"isSuccess"`
	LoginFailedReason string    `gorm:"type:longtext;comment:'로그인 실패 사유'" json:"loginFailedReason"`
	CreatedAt         time.Time `gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt         time.Time `gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt         time.Time `gorm:"type:timestamp;comment:'삭제일자'" json:"deletedAt"`
}

func (l *UserLoginLog) Alias() string {
	return "user_login_log ull"
}

func (l *UserLoginLog) TableName() string {
	return "user_login_log"
}