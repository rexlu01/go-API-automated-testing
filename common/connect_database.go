package common

type MysqlConfig struct {
	Name     string
	Password string
	Host     string
	Port     string
	Database string
}

type RedisConfig struct {
	Network    string
	MasterName string
	RedisAddr  string
	Password   string
	DataDase   int
}
