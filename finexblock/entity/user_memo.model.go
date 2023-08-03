package entity

type UserMemo struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID      uint   `gorm:"comment:'유저 id'" json:"userId"`
	Description string `gorm:"comment:'운영진 메모';type:longtext;not null;" json:"description"`
}

func (m *UserMemo) Alias() string {
	return "user_memo um"
}

func (m *UserMemo) TableName() string {
	return "user_memo"
}