package rules

import (
	"apitool/utils"
)

/**
* 接口响应数据转化成excel数据
* 通过策略模式为不同接口配置不同数据适配方案
 */
func ParseRespData(url string, fileName string, resps []string, apiParams []map[string]string) (exData *utils.ExcelData) {
	//创建Excel数据类型结构体
	//exData = &utils.ExcelData{
	//	ExcelName:       fileName + ".xlsx",
	//	//SheetNameSlice:  sheetNameArr,
	//	//SheetTitleSlice: sheetTitleArr,
	//}
	//
	////策略模式处理，根据url可自动识别匹配解析策略
	//NewParseApiContext(url, fileName, apiParams, resps)

	//sheet和表头解析,通过url获取配置文件参数
	//sheetName := config.Conf.Section(url).Key("sheet_name").Value()
	//sheetTitle := config.Conf.Section(url).Key("sheet_title").Value()
	//sheetTitleStr := splitKeys(sheetTitle, "|")
	//sheetNameArr := splitKeys(sheetName, "|")
	////组装SheetTitleSlice
	//sheetTitleArr := [][]string{}
	//for _, sheetTitle := range sheetTitleStr {
	//	sheetTitleArr = append(sheetTitleArr, splitKeys(sheetTitle, ","))
	//}
	//exData.SheetNameSlice = sheetNameArr
	//exData.SheetTitleSlice = sheetTitleArr
	//
	////excel数据填充处理，需要给不同接口制定对应解析规则
	//SheetDataSlice := parseApi(url, resps)
	//exData.SheetDataSlice = SheetDataSlice
	return
}
