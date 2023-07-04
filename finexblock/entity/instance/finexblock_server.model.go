package instance

type FinexblockServer struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;comment:기본키"`
	Name string `gorm:"type:LONGTEXT;not null;comment:서버명"`
}

func (s *FinexblockServer) Alias() string {
	return "finexblock_server fs"
}

func (s *FinexblockServer) TableName() string {
	return "finexblock_server"
}