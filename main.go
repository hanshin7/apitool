package main

import (
	"apitool/logging"
	"apitool/service"
	. "github.com/sirupsen/logrus"
	"os"
)

var MyLog = logging.MustGetLogger()

func init() {
	//初始化日志配置
	setupLogging()

}

func main() {

	MyLog.Debug("apitool running")
	service.StartService()
}

func setupLogging() {
	SetLevel(DebugLevel)
	logPath := "./MyLog.Log"
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		MyLog.Fatal("Cannot log to file", err.Error())
	}

	//SetFormatter(&TextFormatter{})
	SetOutput(file)
}
