package config

import (
	"github.com/spf13/viper"
	_ "go.uber.org/zap"
)

var (
	Settings ServerConfig
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile("./config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	//给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 传递给全局变量
	Settings = serverConfig
	//color.Blue("11111111", Settings.App.LogsAddress)
}
