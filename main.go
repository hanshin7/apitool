package main

import (
	"apitool/config"
	"apitool/logging"
	"apitool/service"
	"flag"
)

func init() {

	var configPath string
	//参数值地址 参数名称 默认值 参数描述
	flag.StringVar(&configPath, "config", "./static/config.ini", "config path.")
	config.InitConfig(configPath)
	//初始化日志配置
	level := config.Conf.Section("sys").Key("log_level").Value()
	logging.Init(level)
	logging.LogI("配置文件初始化完成")
}

func main() {

	service.StartService()
}
