package utils

import (
	"net/url"
	"testing"
)

func TestHttpPost(t *testing.T) {
	postUrl := "http://139.198.23.208:80/rsj/ent/keyword/query_list"
	params := url.Values{}
	params.Add("key_id", "03136445277747c1")
	params.Add("key_version", "V1.0")
	params.Add("keyword", "阿里巴巴")
	params.Add("page_no", "1")
	params.Add("request_trace_id", "TEST202009291418496")
	params.Add("timestamp", "1601360329")
	params.Add("sign", "4c89567fca7cd23f3defc53b58b9454339d9b14b6ffd3f16ad3e615264773f44")
	resp := httpPost(postUrl, params)
	print(resp)
}

func TestSignByDirectorary(t *testing.T) {
	params := map[string]string{}
	params["keyword"] = "阿里巴巴"
	params["request_trace_id"] = "TEST202009291418496"
	params["key_id"] = "03136445277747c1"
	params["key_version"] = "V1.0"
	params["timestamp"] = "1601360329"
	params["page_no"] = "1"
	params["sign_key"] = "2a6c2430912644fc8258c7be84d3a1ce"
	values := SignByDirectorary(params)
	print(values)
}

func TestRequestApi(t *testing.T) {
	params := map[string]string{}
	params["keyword"] = "阿里巴巴"
	//params["request_trace_id"] = "TEST202009291418496"
	params["key_id"] = "03136445277747c1"
	//params["key_version"] = "V1.0"
	//params["timestamp"] = "1601360329"
	params["page_no"] = "1"
	params["sign_key"] = "2a6c2430912644fc8258c7be84d3a1ce"
	reslut := RequestApi("http://139.198.23.208:80/rsj/ent/keyword/query_list", params)
	println(reslut)
	//循环遍历Map
	//for key, value := range reslut {
	//	fmt.Printf("%s=>%s\n", key, value)
	//}
}

func TestRequestApi1(t *testing.T) {
	params := map[string]string{}
	//params["request_trace_id"] = "TEST202009291418498"
	params["key_id"] = "03136445277747c1"
	//params["key_version"] = "V1.0"
	//params["timestamp"] = "1601360329"
	params["pid"] = "ef7d893e84d6b1b9f41a885be69afb33"
	params["tel"] = "62bd161831fd778434d97324b66daaf0"
	params["sign_key"] = "2a6c2430912644fc8258c7be84d3a1ce"
	reslut := RequestApi("http://139.198.23.208:80/rsj/person/card/finance", params)
	println(reslut)
	//循环遍历Map
	//for key, value := range reslut {
	//	fmt.Printf("%s=>%s\n", key, value)
	//}
}
