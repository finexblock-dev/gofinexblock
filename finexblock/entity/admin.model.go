package entity

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type GradeType string

func (g GradeType) String() string {
	return string(g)
}

func (g GradeType) Validate() error {
	switch g {
	case MAINTAINER, SUPERUSER, SUPPORT:
		return nil
	}
	return errors.New("invalid grade type")
}

const (
	MAINTAINER GradeType = "M"
	SUPERUSER  GradeType = "U"
	SUPPORT    GradeType = "S"
)

type Admin struct {
	ID           uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	Email        string    `json:"email" gorm:"not null;unique;comment:'이메일'"`
	Password     string    `gorm:"comment:'패스워드';not null;type:longtext;" json:"password"`
	Grade        GradeType `json:"grade" gorm:"not null;type:enum('S','M','U');default:'S';comment:'등급';"`
	IsBlocked    bool      `json:"isBlocked" gorm:"comment:'잠김 여부';not null;default:false"`
	InitialLogin bool      `json:"initialLogin" gorm:"default:1;type:tinyint(1);comment:'최초 로그인 여부';not null;"`
	PwdUpdatedAt time.Time `json:"pwdUpdatedAt,omitempty" gorm:"comment:'패스워드 수정일자';type:timestamp;"`
	CreatedAt    time.Time `json:"createdAt" gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;"`
	DeletedAt    time.Time `json:"deletedAt,omitempty" gorm:"comment:'삭제일자';index"`

	ExecuteDeleteLog []AdminDeleteLog `json:"executeDeleteLog" gorm:"foreignKey:ExecutorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	TargetDeleteLog  []AdminDeleteLog `gorm:"foreignKey:TargetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"targetDeleteLog"`

	ExecuteGradeUpdateLog []AdminGradeUpdateLog `gorm:"foreignKey:ExecutorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"executeGradeUpdateLog"`
	TargetGradeUpdateLog  []AdminGradeUpdateLog `gorm:"foreignKey:TargetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"targetGradeUpdateLog"`

	ExecutePasswordLog []AdminPasswordLog `gorm:"foreignKey:ExecutorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"executePasswordLog"`
	TargetPasswordLog  []AdminPasswordLog `gorm:"foreignKey:TargetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"targetPasswordLog"`

	AdminLoginFailedLog []AdminLoginFailedLog `json:"adminLoginFailedLog" gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL"`
	AdminLoginHistory   []AdminLoginHistory   `json:"adminLoginHistory" gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL"`
	AdminAccessToken    []AdminAccessToken    `json:"adminAccessToken" gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL"`
	AdminApiLog         []AdminApiLog         `json:"adminApiLog" gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE;OnDelete:SET NULL"`
}

type PartialAdmin struct {
	ID           uint      `json:"id"`
	Email        string    `json:"email"`
	Grade        GradeType `json:"grade"`
	IsBlocked    bool      `json:"isBlocked"`
	InitialLogin bool      `json:"initialLogin"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	PwdUpdatedAt time.Time `json:"pwdUpdatedAt"`
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