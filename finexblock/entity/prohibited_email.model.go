package entity

import "time"

type ProhibitedEmail struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	Hashed    string    `gorm:"comment:'암호화된 이메일';type:longtext;not null" json:"hashed"`
	CreatedAt time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`
}

func (p *ProhibitedEmail) TableName() string {
	return "prohibited_email"
}