package utils

import "strings"

/**
* url对应的接口名称
 */
func Url2Name(url string) string {
	if strings.Contains(url, "/ent/base/get_ent_info") {
		return "工商照面信息"
	}
	// 添加其他映射
	//...

	return "未知接口"

}
