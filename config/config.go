package config

import (
	"bufio"
	"fmt"
	"gopkg.in/ini.v1"
	"io"
	"os"
	"strings"
)

//var Url2NameMap map[string]string
var Conf *ini.File

func init() {
	//println("初始化加载文件[urlname.conf]")
	////初始化读取url和接口名称映射文件
	//Url2NameMap = initConfig("../Conf/Conf.ini")
	//cfg, err := ini.Load( os.Getenv("GOPATH") + "/config/config.ini")
	//Conf = cfg
	//if err != nil {
	//	fmt.Println("配置文件读取错误", err)
	//	os.Exit(1)
	//}
}
func InitConfig(path string) {
	cfg, err := ini.Load(path)
	Conf = cfg
	if err != nil {
		fmt.Println("配置文件读取错误", err)
		os.Exit(1)
	}
	fmt.Println("配置文件加载完成")
}

/**
* 读取key=value类型的配置文件
 */
func initConfig(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}
