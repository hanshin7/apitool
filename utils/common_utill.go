package utils

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
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
	pList := make([]string, 5)
	pList = append(pList, params["param1"][0])
	pList = append(pList, params["param2"][0])
	pList = append(pList, params["param3"][0])
	pList = append(pList, params["param4"][0])
	pList = append(pList, params["param5"][0])
	for _, p := range pList {
		appendApiParam(p, apiParams)
	}
}

func appendApiParam(param string, apiParams map[string]string) {
	if strings.Contains(param, ":") {
		arr := strings.Split(param, ":")
		apiParams[arr[0]] = arr[1]
	}
}

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
