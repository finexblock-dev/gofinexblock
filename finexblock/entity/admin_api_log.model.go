package entity

import (
	"errors"
	"time"
)

type ApiMethod string

func (a ApiMethod) String() string {
	return string(a)
}

func (a ApiMethod) Validate() error {
	switch a {
	case GetMethod, PostMethod, PutMethod, PatchMethod, DeleteMethod:
		return nil
	}
	return errors.New("invalid api method")
}

const (
	GetMethod    ApiMethod = "GET"
	PostMethod   ApiMethod = "POST"
	PutMethod    ApiMethod = "PUT"
	PatchMethod  ApiMethod = "PATCH"
	DeleteMethod ApiMethod = "DELETE"
)

type AdminApiLog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	AdminID   uint      `gorm:"comment:'운영진 id'" json:"adminId"`
	IP        string    `gorm:"type:longtext;not null;comment:'IP 주소'" json:"ip"`
	Endpoint  string    `gorm:"type:longtext;not null;comment:'API 엔드포인트'" json:"endpoint"`
	Method    ApiMethod `gorm:"type:enum('GET', 'POST', 'PUT', 'PATCH', 'DELETE');not null;comment:'메소드'" json:"method"`
	CreatedAt time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`
}

func (a *AdminApiLog) Alias() string {
	return "admin_api_log apl"
}

func (a *AdminApiLog) TableName() string {
	return "admin_api_log"
}