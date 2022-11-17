package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	fmt.Println("config")
	// https://github.com/spf13/viper#reading-config-files
	viper.SetConfigName("app")      // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // path to look for the config file in
	err := viper.ReadInConfig()
	if err != nil {
		panic("讀取設定檔出現錯誤，原因為：" + err.Error())
	}

	fmt.Println("application port = " + viper.GetString("application.port"))
}
