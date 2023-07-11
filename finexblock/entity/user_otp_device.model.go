package entity

import "time"

type UserOtpDevice struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'"`
	UserID    uint      `gorm:"comment:'유저 id'"`
	Secret    string    `gorm:"comment:'시크릿'"`
	QrImage   string    `gorm:"comment:'QR 이미지'"`
	CreatedAt time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt time.Time `gorm:"comment:'삭제일자';index"`
}

func (o *UserOtpDevice) Alias() string {
	return "user_otp_device uod"
}

func (o *UserOtpDevice) TableName() string {
	return "user_otp_device"
}
