package admin

import (
	"time"
)

type AdminLoginHistory struct {
	ID       uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	AdminID  uint      `json:"admin_id,omitempty" gorm:"comment:'운영진 id'"`
	LoggedAt time.Time `json:"logged_at" gorm:"comment:'로그인 일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (l *AdminLoginHistory) Alias() string {
	return "admin_login_history alh"
}

func (l *AdminLoginHistory) TableName() string {
	return "admin_login_history"
}
