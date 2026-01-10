package utils

var DB *gorm.DB

func InitConfig() {
	common.VP = viper.GetViper();
	common.VP.SetCOnfigName("app")
	common.VP.AddConfigPath("config")
	return common.VP.ReadInConfig()
}


func InitMySQL() {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel: logger.Info,
		Colorful: true,
	})

	mysqlConfig, err := config.GetMysqlConfig()
	if err != nil {
		panic("failed to read mysqlConfig")
	}

	common.DB, err = gorm.Open(mysql.Open(mysql.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		mysqlConfig.User,
		mysqlConfig.PassWord,
		mysqlConfig.Ip,
		mysqlConfig.Port,
		mysqlConfig.Database)),
		&gorm.Config{logger: newLogger})

	if err != nil {
		panic("failed to connect database")
	}
	common.DB.AutoMigrate(&models.UserBasic{})
}




