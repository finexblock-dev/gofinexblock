package admin

import (
	"time"
)

type AdminApiLog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	AdminID   uint      `json:"admin_id,omitempty" gorm:"comment:'운영진 id'"`
	IP        string    `json:"ip,omitempty" gorm:"type:longtext;not null;comment:'IP 주소'"`
	Endpoint  string    `json:"endpoint,omitempty" gorm:"type:longtext;not null;comment:'API 엔드포인트'"`
	Method    string    `json:"method,omitempty" gorm:"type:enum('GET', 'POST', 'PUT', 'PATCH', 'DELETE');not null;comment:'메소드'"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (a *AdminApiLog) Alias() string {
	return "admin_api_log apl"
}

func (a *AdminApiLog) TableName() string {
	return "admin_api_log"
}
