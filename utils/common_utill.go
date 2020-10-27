package utils

import (
	"apitool/config"
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
* url对应的接口名称
 */
func GetNameByurl(url string) string {
	for k, v := range config.Url2NameMap {
		if strings.Contains(url, k) {
			return v
		}
	}
	return "未知接口"
}

//func WriteFile(file string , lines *[]string) {
//	for _,line := range *lines {
//		ioutil.WriteFile(file, []byte(line + "\n"), 0666)
//	}
//
//}

func WriteFile(filePath string, lines *[]string) error {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("create file error: %v\n", err)
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, v := range *lines {
		fmt.Fprintln(w, v)
	}
	return w.Flush()
}
