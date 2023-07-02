package user

import "time"

type UserPaypalAccount struct {
	ID            uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID        uint      `json:"user_id,omitempty" gorm:"comment:'유저 id'"`
	PaypalAddress string    `json:"paypal_address,omitempty" gorm:"comment:'페이팔 주소';not null;type:longtext"`
	CreatedAt     time.Time `json:"created_at" gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt     time.Time `json:"deleted_at" gorm:"comment:'삭제일자';index"`
}

func (p *UserPaypalAccount) Alias() string {
	return "user_paypal_account upa"
}

func (p *UserPaypalAccount) TableName() string {
	return "user_paypal_account"
}
