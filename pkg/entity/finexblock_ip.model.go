package entity

type FinexblockServerIP struct {
	ID       uint   `gorm:"column:id;primary_key;auto_increment;comment:'기본키'" json:"Id"`
	ServerID uint   `gorm:"column:server_id;not null;comment:'서버 id'" json:"serverId"`
	IP       string `gorm:"column:ip;type:longtext;not null;comment:'IP 주소'" json:"ip"`
}

// TableName sets the insert table name for this structs type
func (f *FinexblockServerIP) TableName() string {
	return "finexblock_server_ip"
}

// Alias sets the alias name for this structs type
func (f *FinexblockServerIP) Alias() string {
	return "finexblock_server_ip fsi"
}
