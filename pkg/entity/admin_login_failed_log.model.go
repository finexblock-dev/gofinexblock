package entity

import (
	"time"
)

type AdminLoginFailedLog struct {
	ID       uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	AdminID  uint      `json:"adminId" gorm:"comment:'운영진 id'"`
	FailedAt time.Time `json:"failedAt" gorm:"comment:'로그인 실패 일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (l *AdminLoginFailedLog) Alias() string {
	return "admin_login_failed_log alfl"
}

func (l *AdminLoginFailedLog) TableName() string {
	return "admin_login_failed_log"
}