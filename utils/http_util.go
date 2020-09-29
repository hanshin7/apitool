package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

/*
 * api调用函数，返回结果map
 */
func requestApi(url string, params map[string]string) map[string]interface{} {
	values := SignByDirectorary(params)
	resultStr := HttpPost(url, values)
	print(resultStr)
	//json结果字符串转map对象
	resultMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(resultStr), &resultMap)
	if err == nil {
		return nil
	}
	return resultMap
}

/**
 * 发送POST请求
 * url:请求地址
 * params:POST请求提交的数据
 */
func HttpPost(url string, values url.Values) string {
	resp, err := http.PostForm(url, values)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

/**
 * 请求参数签名，字典序
 */
func SignByDirectorary(params map[string]string) url.Values {
	//存放参数map中的所有key
	keyList := []string{}
	for k, _ := range params {
		keyList = append(keyList, k)
	}
	//对key进行排序
	sort.Strings(keyList)
	buf := bytes.Buffer{}
	for _, v := range keyList {
		if v != "sign_key" {
			buf.WriteString(v + "=" + params[v] + "&")
		}
	}
	buf.WriteString("sign_key=" + params["sign_key"])
	//println(buf.String())
	sha256Val := sha256.Sum256(buf.Bytes())
	params["sign"] = fmt.Sprintf("%x", sha256Val)
	values := url.Values{}
	for k, v := range params {
		if k != "sign_key" {
			values.Add(k, v)
		}
	}
	return values
}
