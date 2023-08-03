package entity

import (
	"time"
)

type AdminLoginHistory struct {
	ID       uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	AdminID  uint      `json:"adminId,omitempty" gorm:"comment:'운영진 id'"`
	LoggedAt time.Time `json:"loggedAt" gorm:"comment:'로그인 일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (l *AdminLoginHistory) Alias() string {
	return "admin_login_history alh"
}

func (l *AdminLoginHistory) TableName() string {
	return "admin_login_history"
}