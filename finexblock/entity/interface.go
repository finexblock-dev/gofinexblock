package finexblock_server

type Model interface {
	TableName() string
	Alias() string
}
