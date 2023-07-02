package admin

import (
	"time"
)

type AdminAccessToken struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	AdminID   uint      `json:"admin_id,omitempty" gorm:"comment:'운영진 id'"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	ExpiredAt time.Time `json:"expired_at" gorm:"comment:'만료일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (a *AdminAccessToken) Alias() string {
	return "admin_access_token aat"
}

func (a *AdminAccessToken) TableName() string {
	return "admin_access_token"
}
