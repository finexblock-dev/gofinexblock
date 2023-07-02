package admin

import (
	"time"
)

type AdminGradeUpdateLog struct {
	ID         uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	ExecutorID uint      `json:"executor_id,omitempty" gorm:"comment:'변경한 운영진 id'"`
	TargetID   uint      `json:"target_id,omitempty" gorm:"comment:'변경된 운영진 id'"`
	PrevGrade  string    `json:"prev_grade,omitempty" gorm:"not null;type:enum('S','M','N');comment:'수정 전 등급';"`
	CurrGrade  string    `json:"curr_grade,omitempty" gorm:"not null;type:enum('S','M','N');comment:'수정 후 등급';"`
	CreatedAt  time.Time `json:"created_at" gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`

	//Executor Admin `gorm:"foreignKey:ExecutorID;references:ID;constraint:OnUpdate:CASCADE"`
	//Target   Admin `gorm:"foreignKey:TargetID;references:ID;constraint:OnUpdate:CASCADE"`
}

func (g *AdminGradeUpdateLog) Alias() string {
	return "admin_grade_update_log agul"
}

func (g *AdminGradeUpdateLog) TableName() string {
	return "admin_grade_update_log"
}
