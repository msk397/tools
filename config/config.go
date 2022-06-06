package config

type ServerConfig struct {
	App    appConfig    `structure:"app"`
	Mysql  MysqlConfig  `structure:"mysql"`
	Notion notionConfig `structure:"notion"`
}

type appConfig struct {
	Name        string `structure:"name"`
	Version     string `structure:"version"`
	Ip          string `structure:"ip"`
	Port        string `structure:"port"`
	Env         string `structure:"env"`
	LogsAddress string `structure:"logsAddress"`
}

type MysqlConfig struct {
	Host     string `structure:"host"`
	Port     string `structure:"port"`
	Name     string `structure:"name"`
	Password string `structure:"password"`
	DBName   string `structure:"dbName"`
}

type notionConfig struct {
	Auth          string `structure:"auth"`
	CryptoBlockId string `structure:"cryptoBlockId"`
	Version       string `structure:"version"`
}
