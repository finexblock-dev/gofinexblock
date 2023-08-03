package entity

type FinexblockServer struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;comment:기본키" json:"id"`
	Name string `gorm:"type:LONGTEXT;not null;comment:서버명" json:"name"`
}

func (s *FinexblockServer) Alias() string {
	return "finexblock_server fs"
}

func (s *FinexblockServer) TableName() string {
	return "finexblock_server"
}