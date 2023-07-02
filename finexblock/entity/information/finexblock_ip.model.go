package information

import "gorm.io/gorm"

type FinexblockServerIP struct {
	ID       uint64 `gorm:"column:id;primary_key;auto_increment;comment:'기본키'"`
	ServerID uint64 `gorm:"column:server_id;not null;comment:'서버 id'"`
	IP       string `gorm:"column:ip;type:longtext;not null;comment:'IP 주소'"`
	gorm.Model
}

// TableName sets the insert table name for this struct type
func (f *FinexblockServerIP) TableName() string {
	return "finexblock_server_ip"
}

// Alias sets the alias name for this struct type
func (f *FinexblockServerIP) Alias() string {
	return "finexblock_server_ip fsi"
}
