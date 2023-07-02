package user

type UserMemo struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id,omitempty"`
	UserID      uint   `json:"user_id,omitempty" gorm:"comment:'유저 id'"`
	Description string `json:"description,omitempty" gorm:"comment:'운영진 메모';type:longtext;not null;"`
}

func (m *UserMemo) Alias() string {
	return "user_memo um"
}

func (m *UserMemo) TableName() string {
	return "user_memo"
}
