package utils

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

/**
* url对应的接口名称
 */
//func GetNameByurl(url string) string {
//	for k, v := range config.Url2NameMap {
//		if strings.Contains(url, k) {
//			return v
//		}
//	}
//	return "未知接口"
//}

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

func ParseFormParams(params url.Values, apiParams map[string]string) {
	param1 := params["param1"][0]
	param2 := params["param2"][0]
	param3 := params["param3"][0]
	param4 := params["param4"][0]
	param5 := params["param5"][0]

	if strings.Contains(param1, ":") {
		arr := strings.Split(param1, ":")
		apiParams[arr[0]] = arr[1]
	}
	if strings.Contains(param2, ":") {
		arr := strings.Split(param2, ":")
		apiParams[arr[0]] = arr[1]
	}
	if strings.Contains(param3, ":") {
		arr := strings.Split(param3, ":")
		apiParams[arr[0]] = arr[1]
	}
	if strings.Contains(param4, ":") {
		arr := strings.Split(param4, ":")
		apiParams[arr[0]] = arr[1]
	}
	if strings.Contains(param5, ":") {
		arr := strings.Split(param5, ":")
		apiParams[arr[0]] = arr[1]
	}
}
