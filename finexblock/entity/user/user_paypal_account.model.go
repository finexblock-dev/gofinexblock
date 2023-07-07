package user

import "time"

type UserPaypalAccount struct {
	ID            uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	UserID        uint      `gorm:"comment:'유저 id'"`
	PaypalAddress string    `gorm:"comment:'페이팔 주소';not null;type:longtext"`
	CreatedAt     time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt     time.Time `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt     time.Time `gorm:"comment:'삭제일자';index"`
}

func (p *UserPaypalAccount) Alias() string {
	return "user_paypal_account upa"
}

func (p *UserPaypalAccount) TableName() string {
	return "user_paypal_account"
}
