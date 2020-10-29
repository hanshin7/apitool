package rules

import (
	"apitool/config"
	"apitool/model"
	"apitool/utils"
	"encoding/json"
	"strings"
)

/**
* 个人五要素接口
 */

type ICApi struct {
	apiParams []map[string]string
}

func newICApi(apiParams []map[string]string) ICApi {
	instance := new(ICApi)
	instance.apiParams = apiParams
	return *instance
}

/**
* 单个sheet，接口响应数据为单层,不存在嵌套类型,动态指定接口返回字段
 */
func (obj ICApi) ParseApi(url string, resps []string, exData *utils.ExcelData) {
	//sheet和表头解析,通过url获取配置文件参数
	sheetName := config.Conf.Section(url).Key("sheet_name").Value()
	sheetNameArr := splitKeys(sheetName, "|")
	sheetTitle := config.Conf.Section(url).Key("sheet_title_map").Value()
	sheetTitleStr := splitKeys(sheetTitle, "|")
	sheetTitleMap := make(map[string]string)
	//将配置文件中kv映射存到map
	for _, m := range sheetTitleStr {
		kvArr := strings.Split(m, ",")
		sheetTitleMap[kvArr[0]] = kvArr[1]
	}

	//组装SheetTitleSlice,通过请求参数参数字段组装,取第一条参数即可
	outstyle := obj.apiParams[0]["outstyle"]
	//去掉|分隔后面内容
	outstyle = strings.Split(outstyle, "|")[0]
	sheetTitleArr := [][]string{}
	//英文字段替换成中文描述
	titleReplArr := []string{"接口响应结果", "名称", "身份证号"}
	colArr := splitKeys(outstyle, ",")
	for _, v := range splitKeys(outstyle, ",") {
		val := sheetTitleMap[v]
		titleReplArr = append(titleReplArr, val)
	}
	sheetTitleArr = append(sheetTitleArr, titleReplArr)
	exData.SheetNameSlice = sheetNameArr
	exData.SheetTitleSlice = sheetTitleArr

	//excel数据填充处理，需要给不同接口制定对应解析规则
	SheetDataSlice := parseIcData(url, obj.apiParams, resps, &colArr)
	exData.SheetDataSlice = SheetDataSlice
}

/**
* 单sheet接口解析,data数据根据不同类型不同处理
 */
func parseIcData(url string, apiParams []map[string]string, resps []string, colArr *[]string) [][][]string {
	//对应sheet
	sheetSlice := [][]string{}
	for k, resp := range resps {
		var apiResp model.ApiRespMsg
		err := json.Unmarshal([]byte(resp), &apiResp)
		if err != nil {
			panic(err)
		}

		var lineSlice []string
		//其他列根据配置文件中的key顺序取值
		if apiResp.Data == nil {
			lineSlice = parseIcLineData(url, map[string]interface{}{}, colArr)
		} else {
			lineSlice = parseIcLineData(url, apiResp.Data.(map[string]interface{}), colArr)
		}
		//第一列固定为接口响应结果信息
		lineSlice[0] = apiResp.Code + "|" + apiResp.Msg
		lineSlice[1] = apiParams[k]["name"]
		lineSlice[2] = apiParams[k]["pid"]
		sheetSlice = append(sheetSlice, lineSlice)
	}
	sheetDataSlice := [][][]string{sheetSlice}
	return sheetDataSlice
}

func parseIcLineData(url string, data map[string]interface{}, colArr *[]string) (lineSlice []string) {

	keyArr := *colArr
	//sheet中一行数据,接口字段列数加上额外一列结果数据
	lineSlice = make([]string, len(keyArr)+3)
	for i, key := range keyArr {
		value, exist := data[key]
		//第0位保留给额外附件字段
		if exist {
			lineSlice[i+1] = value.(string)
		} else {
			lineSlice[i+1] = ""
		}
	}
	return lineSlice
}
