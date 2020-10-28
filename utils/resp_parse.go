package utils

import (
	"apitool/config"
	"apitool/model"
	"encoding/json"
	"strings"
)

//接口响应数据转化成excel数据
func ParseRespData(url string, fileName string, resps *[]string) (exData *ExcelData) {
	//sheet和表头解析,通过url获取配置文件参数
	sheetName := config.Conf.Section(url).Key("sheet_name").Value()
	sheetTitle := config.Conf.Section(url).Key("sheet_title").Value()
	sheetTitleStr := splitKeys(sheetTitle, "|")
	//组装SheetTitleSlice
	sheetTitleArr := [][]string{}
	for _, sheetTitle := range sheetTitleStr {
		sheetTitleArr = append(sheetTitleArr, splitKeys(sheetTitle, ","))
	}

	//创建Excel数据类型结构体
	exData = &ExcelData{
		ExcelName:       fileName + ".xlsx",
		SheetNameSlice:  splitKeys(sheetName, "|"),
		SheetTitleSlice: sheetTitleArr,
	}

	//excel数据填充处理，需要给不同接口制定对应解析规则
	SheetDataSlice := parseApi(url, resps)
	exData.SheetDataSlice = SheetDataSlice
	return
}

/**
* 单sheet接口解析,data数据根据不同类型不同处理
 */
func parseApi(url string, resps *[]string) [][][]string {
	//对应sheet
	sheetSlice := [][]string{}
	for _, resp := range *resps {
		var apiResp model.ApiRespMsg
		err := json.Unmarshal([]byte(resp), &apiResp)
		if err != nil {
			panic(err)
		}

		var lineSlice []string
		//其他列根据配置文件中的key顺序取值
		if apiResp.Data == nil {
			lineSlice = parseSimpleData(url, map[string]interface{}{})
		} else {
			lineSlice = parseSimpleData(url, apiResp.Data.(map[string]interface{}))
		}
		//第一列固定为接口响应结果信息
		lineSlice[0] = apiResp.Code + "|" + apiResp.Msg
		sheetSlice = append(sheetSlice, lineSlice)
	}
	sheetDataSlice := [][][]string{sheetSlice}
	return sheetDataSlice
}

func parseSimpleData(url string, data map[string]interface{}) (lineSlice []string) {
	sheetKey := config.Conf.Section(url).Key("sheet_key").Value()
	//取第一个'|'前数据
	keystr := splitKeys(sheetKey, "|")[0]
	keyArr := splitKeys(keystr, ",")
	//sheet中一行数据,接口字段列数加上额外一列结果数据
	lineSlice = make([]string, len(keyArr)+1)
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

func splitKeys(str string, sep string) []string {
	arr := strings.Split(str, sep)
	if arr[len(arr)-1] == "" {
		return arr[0 : len(arr)-1]
	}
	return arr
}
