package utils

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("ginchat\\config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))
}


func INitMySQL() {
	db, err := gorm.Open(mysql.Open(""), &gorn.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}




