package entity

import (
	"time"
)

type AdminPasswordLog struct {
	ID         uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	ExecutorID uint      `json:"executor_id,omitempty" gorm:"comment:'변경한 운영진 id'"`
	TargetID   uint      `json:"target_id,omitempty" gorm:"comment:'변경된 운영진 id'"`
	CreatedAt  time.Time `json:"created_at" gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`

	//Executor Admin `gorm:"foreignKey:ExecutorID;references:ID;constraint:OnUpdate:CASCADE"`
	//Target   Admin `gorm:"foreignKey:TargetID;references:ID;constraint:OnUpdate:CASCADE"`
}

func (p *AdminPasswordLog) Alias() string {
	return "admin_password_log apl"
}

func (p *AdminPasswordLog) TableName() string {
	return "admin_password_log"
}
