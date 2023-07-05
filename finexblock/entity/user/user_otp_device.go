package user

import "time"

type UserOtpDevice struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID    uint      `json:"user_id" gorm:"comment:'유저 id'"`
	Secret    string    `json:"secret" gorm:"comment:'시크릿'"`
	QrImage   string    `json:"qrImage" gorm:"comment:'QR 이미지'"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	DeletedAt time.Time `json:"deleted_at" gorm:"comment:'삭제일자';index"`
}

func (o *UserOtpDevice) Alias() string {
	return "user_otp_device uod"
}

func (o *UserOtpDevice) TableName() string {
	return "user_otp_device"
}
