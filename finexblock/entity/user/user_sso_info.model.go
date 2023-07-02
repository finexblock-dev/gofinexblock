package user

import "time"

type UserSingleSignOnInfo struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID    uint      `json:"user_id,omitempty" gorm:"comment:'유저 id'"`
	SSOType   string    `json:"sso_type,omitempty" gorm:"comment:'SSO 타입';type:enum('APPLE', 'GOOGLE', 'METAVERSE');not null"`
	AuthCode  string    `json:"auth_code,omitempty" gorm:"not null;comment:'코드';"`
	Email     string    `json:"email,omitempty" gorm:"size:100;comment:'이메일';not null"`
	ExpiredAt time.Time `json:"expired_at" gorm:"type:timestamp;comment:'만료일자';not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" gorm:"type:timestamp;comment:'삭제일자'"`
}

func (S *UserSingleSignOnInfo) Alias() string {
	return "user_sso_info usi"
}

func (S *UserSingleSignOnInfo) TableName() string {
	return "user_sso_info"
}
