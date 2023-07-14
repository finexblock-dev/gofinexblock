package entity

import (
	"time"
)

type ApiMethod string

const (
	GetMethod    ApiMethod = "GET"
	PostMethod   ApiMethod = "POST"
	PutMethod    ApiMethod = "PUT"
	PatchMethod  ApiMethod = "PATCH"
	DeleteMethod ApiMethod = "DELETE"
)

type AdminApiLog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	AdminID   uint      `gorm:"comment:'운영진 id'"`
	IP        string    `gorm:"type:longtext;not null;comment:'IP 주소'"`
	Endpoint  string    `gorm:"type:longtext;not null;comment:'API 엔드포인트'"`
	Method    ApiMethod `gorm:"type:enum('GET', 'POST', 'PUT', 'PATCH', 'DELETE');not null;comment:'메소드'"`
	CreatedAt time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (a *AdminApiLog) Alias() string {
	return "admin_api_log apl"
}

func (a *AdminApiLog) TableName() string {
	return "admin_api_log"
}
