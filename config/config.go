package config

type Mysql struct {
	Ip string
	Port string
	Database string
	User string
	PassWord string
}

func GetMysqlConfig() (*Mysql, error) {
	mapConfig := ReadSection("mysql")
	mysqlConfig := Mysql{}
	mysqlConfig.Ip = mapConfig["ip"]
	mysqlCOnfig.Port = mapConfig["port"]
	mysqlConfig.Database = mapConfig["database"]
	mysqlConfig.User = mapConfig["user"]
	mysqlConfig.PassWord = mapConfig["password"]
	return &mysqlConfig, nil
}

func ReadSection(name string) map[string]string {
	return common.VP.GetStringMapString(name)
}