package rules

import "apitool/utils"

/**
* 定义策略接口
 */
type parseApiInterface interface {
	//函数返回封装好的数据体引用地址
	ParseApi(url string, resps []string, exData *utils.ExcelData)
}

//用来做策略筛选
type parseApiStrategy struct {
	Strategy parseApiInterface
}

//外部调用策略的入口
func NewParseApiStrategy(url string, fileName string, apiParams []map[string]string, resps []string) (exData *utils.ExcelData) {
	//创建Excel数据类型结构体,所有接口公用
	exData = &utils.ExcelData{
		ExcelName: fileName + ".xlsx",
		//SheetNameSlice:  sheetNameArr,
		//SheetTitleSlice: sheetTitleArr,
	}

	p := new(parseApiStrategy)
	switch url {
	case "rsj/person/railwaylabel/query":
		p.Strategy = newICApi(apiParams)
	case "rsj/person/gsinfo/query":
		p.Strategy = newUsualChildrenApi(apiParams)
	case "rsj/ent/base/get_ent_check":
		p.Strategy = newUsualChildrenApi(apiParams)
	case "rsj/ent/base/get_ent_info":
		p.Strategy = newUsualChildrenApi(apiParams)
	default:
		//普通接口解析
		p.Strategy = newUsualApi()

	}
	//调用真正的执行方法,参数为每个接口通用参数
	p.Strategy.ParseApi(url, resps, exData)
	return
}
