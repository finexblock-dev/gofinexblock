package user

import "time"

type ProhibitedEmail struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	Hashed    string    `gorm:"comment:'암호화된 이메일';type:longtext;not null"`
	CreatedAt time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
}

func (p *ProhibitedEmail) TableName() string {
	return "prohibited_email"
}
