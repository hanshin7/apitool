package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"time"
)

/*
 * api调用函数，返回结果map
 */
func RequestApi(url string, apiParams map[string]string) map[string]interface{} {
	//此处检测空，在调用处约束
	key_id := apiParams["key_id"]
	sign_key := apiParams["sign_key"]
	params := map[string]string{}
	//公共参数赋值
	stampStr := fmt.Sprint(time.Now().Unix())
	params["request_trace_id"] = "TEST" + time.Now().Format("20060102") + stampStr + fmt.Sprint(rand.Intn(1000))
	params["key_id"] = key_id
	params["key_version"] = "V1.0"
	params["timestamp"] = stampStr
	params["sign_key"] = sign_key
	//添加api需要的kv
	for k, v := range apiParams {
		params[k] = v
	}

	values := SignByDirectorary(params)
	resultStr := httpPost(url, values)
	println(resultStr)
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
func httpPost(url string, values url.Values) string {
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

/*时间戳->字符串*/
func Stamp2Str(stamp int64) string {
	timeLayout := "2006-01-02 15:04:05"
	str := time.Unix(stamp/1000, 0).Format(timeLayout)
	return str
}