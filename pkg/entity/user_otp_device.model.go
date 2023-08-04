package entity

import "time"

type UserOtpDevice struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID    uint      `gorm:"comment:'유저 id'" json:"userId"`
	Secret    string    `gorm:"comment:'시크릿'" json:"secret"`
	QrImage   string    `gorm:"comment:'QR 이미지'" json:"QRImage"`
	CreatedAt time.Time `gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp" json:"updatedAt"`
	DeletedAt time.Time `gorm:"comment:'삭제일자';index" json:"deletedAt"`
}

func (o *UserOtpDevice) Alias() string {
	return "user_otp_device uod"
}

func (o *UserOtpDevice) TableName() string {
	return "user_otp_device"
}