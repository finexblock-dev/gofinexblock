package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type GradeType string

const (
	MAINTAINER GradeType = "MAINTAINER"
	SUPERUSER  GradeType = "SUPERUSER"
	SUPPORT    GradeType = "SUPPORT"
)

type Admin struct {
	ID           uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id,omitempty"`
	Email        string    `json:"email,omitempty" gorm:"not null;unique;comment:'이메일'"`
	Password     string    `gorm:"comment:'패스워드';not null;type:longtext;" json:"password,omitempty"`
	Grade        GradeType `json:"grade,omitempty" gorm:"not null;type:enum('S','M','U');default:'S';comment:'등급';"`
	IsBlocked    bool      `json:"is_blocked" gorm:"comment:'잠김 여부';not null;default:false"`
	InitialLogin bool      `json:"initial_login" gorm:"default:1;type:tinyint(1);comment:'최초 로그인 여부';not null;"`
	PwdUpdatedAt time.Time `json:"pwd_updated_at" gorm:"comment:'패스워드 수정일자';type:timestamp;"`
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;"`
	DeletedAt    time.Time `json:"deleted_at,omitempty" gorm:"comment:'삭제일자';index"`

	ExecuteDeleteLog []AdminDeleteLog `json:"execute_delete_log,omitempty" gorm:"foreignKey:ExecutorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	TargetDeleteLog  []AdminDeleteLog `gorm:"foreignKey:TargetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"target_delete_log,omitempty"`

	ExecuteGradeUpdateLog []AdminGradeUpdateLog `gorm:"foreignKey:ExecutorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"execute_grade_update_log,omitempty"`
	TargetGradeUpdateLog  []AdminGradeUpdateLog `gorm:"foreignKey:TargetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"target_grade_update_log,omitempty"`

	ExecutePasswordLog []AdminPasswordLog `gorm:"foreignKey:ExecutorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"execute_password_log,omitempty"`
	TargetPasswordLog  []AdminPasswordLog `gorm:"foreignKey:TargetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"target_password_log,omitempty"`

	AdminLoginFailedLog []AdminLoginFailedLog `json:"admin_login_failed_log,omitempty" gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL"`
	AdminLoginHistory   []AdminLoginHistory   `json:"admin_login_history,omitempty" gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL"`
	AdminAccessToken    []AdminAccessToken    `json:"admin_access_token,omitempty" gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL"`
	AdminApiLog         []AdminApiLog         `json:"admin_api_log,omitempty" gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL"`
}

type PartialAdmin struct {
	ID           uint      `json:"id"`
	Email        string    `json:"email"`
	Grade        string    `json:"grade"`
	IsBlocked    bool      `json:"is_blocked"`
	InitialLogin bool      `json:"initial_login"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	PwdUpdatedAt time.Time `json:"pwd_updated_at"`
}

func (a *Admin) Alias() string {
	return "admin a"
}

func (a *Admin) TableName() string {
	return "admin"
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)
	return nil
}

func (a *Admin) BeforeUpdate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)
	return nil
}
