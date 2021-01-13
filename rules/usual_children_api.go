package rules

/**
* 带一层子节点通用api解析策略
*
 */
import (
	"apitool/config"
	"apitool/model"
	"apitool/utils"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type UsualChildrenApi struct {
	apiParams []map[string]string
}

func newUsualChildrenApi(apiParams []map[string]string) UsualChildrenApi {
	instance := new(UsualChildrenApi)
	instance.apiParams = apiParams
	return *instance
}

/**
* 多个sheet
 */
func (obj UsualChildrenApi) ParseApi(url string, resps []string, exData *utils.ExcelData) {
	//sheet和表头解析,通过url获取配置文件参数
	sheetName := config.Conf.Section(url).Key("sheet_name").Value()
	sheetNameArr := splitKeys(sheetName, "|")
	sheetNameKey := config.Conf.Section(url).Key("sheet_name_key").Value()
	sheetNameKeyArr := splitKeys(sheetNameKey, "|")

	//公共列，数据字段位于data下第一层
	baseColKey := config.Conf.Section(url).Key("base_col_key").Value()
	baseColKeyArr := splitKeys(baseColKey, ",")

	//组装SheetTitleSlice
	allSheetTitleArr := [][]string{}
	//每个sheet的列的英文字段名
	allSheetTitleKeyArr := [][]string{}
	for _, v := range sheetNameKeyArr {
		sheetTitle := config.Conf.Section(url).Key("sheet_title_" + v).Value()
		//当前sheet的列切片
		sheetTitleArr := splitKeys(sheetTitle, ",")
		allSheetTitleArr = append(allSheetTitleArr, sheetTitleArr)

		sheetTitleKey := config.Conf.Section(url).Key("sheet_title_key_" + v).Value()
		sheetTitleKeyArr := splitKeys(sheetTitleKey, ",")
		allSheetTitleKeyArr = append(allSheetTitleKeyArr, sheetTitleKeyArr)
	}

	exData.SheetNameSlice = sheetNameArr
	exData.SheetTitleSlice = allSheetTitleArr

	//excel数据填充处理，需要给不同接口制定对应解析规则
	SheetDataSlice := parseUsualChildrenData(url, obj.apiParams, resps, sheetNameKeyArr, allSheetTitleKeyArr, baseColKeyArr)
	exData.SheetDataSlice = SheetDataSlice

}

/**
* 多sheet接口解析,最外层解析
 */
func parseUsualChildrenData(url string, apiParams []map[string]string, resps []string, sheetNameKeyArr []string, allSheetTitleKeyArr [][]string, baseColKeyArr []string) [][][]string {
	//所有sheet数据集合
	sheetDataSlice := [][][]string{}

	for i, resp := range resps {
		var apiResp model.ApiRespMsg
		err := json.Unmarshal([]byte(resp), &apiResp)
		if err != nil {
			panic(err)
		}

		//遍历每个sheet，取接口值
		for k, v := range sheetNameKeyArr {
			//单个sheet的数据
			//var sheetSlice [][]string
			//其他列根据配置文件中的key顺序取值
			if apiResp.Data == nil {
				//sheetSlice = parseChildrenData(url, map[string]interface{}{}, v, allSheetTitleKeyArr[k])
			} else {
				sheetSlice := parseChildrenData(url, apiParams[i], apiResp.Data.(map[string]interface{}), v, allSheetTitleKeyArr[k], baseColKeyArr)
				//切片合并
				if len(sheetDataSlice) == 0 {

				}
				//sheetDataSlice[k] = append(sheetDataSlice[k], sheetSlice...)
				sheetDataSlice = append(sheetDataSlice, sheetSlice)
			}
		}

		//var lineSlice []string
		////其他列根据配置文件中的key顺序取值
		//if apiResp.Data == nil {
		//	lineSlice = parseChildrenData(url, map[string]interface{}{}, sheetNameKeyArr, allSheetTitleKeyArr)
		//} else {
		//	lineSlice = parseChildrenData(url, apiResp.Data.(map[string]interface{}), sheetNameKeyArr, allSheetTitleKeyArr)
		//}
		////第一列固定为接口响应结果信息
		//lineSlice[0] = apiResp.Code + "|" + apiResp.Msg
		//sheetSlice = append(sheetSlice, lineSlice)
	}
	//sheetDataSlice := [][][]string{sheetSlice}
	return sheetDataSlice
}

/**
* 单个sheet取值
 */
func parseChildrenData(url string, apiParam map[string]string, data map[string]interface{}, sheetNameKey string, sheetTitleKeyArr []string, baseColKeyArr []string) (sheetSlice [][]string) {

	//var baseColLen int
	//if baseColKeyArr == nil {
	//	baseColLen = 0
	//} else {
	//	baseColLen = len(baseColKeyArr)
	//}

	baseColVal := []string{}
	for _, v := range baseColKeyArr {
		value, exist := data[v]
		if exist {
			baseColVal = append(baseColVal, value.(string))
		} else {
			//特殊接口特色处理 从请求参数获取
			if strings.EqualFold(url, "rsj/person/gsinfo/query") ||
				strings.EqualFold(url, "rsj/person/gsinfo/common") {
				baseColVal = append(baseColVal, apiParam[v])
			} else {
				baseColVal = append(baseColVal, "")
			}
		}
	}
	value, exist := data[sheetNameKey]
	if exist {
		//切片，切片存储类型为map
		var lineSlice []map[string]interface{}
		valType := fmt.Sprintf("%v", reflect.TypeOf(value))
		//为了统一格式处理，map先存放到数组
		if strings.EqualFold(valType, "map[string]interface {}") {
			v := value.(map[string]interface{})
			lineSlice = append(lineSlice, v)
		} else {
			//不为空时赋值
			if strings.EqualFold(valType, "[]interface {}") {
				line := value.([]interface{})
				if len(line) > 0 {
					for _, l := range line {
						v := l.(map[string]interface{})
						lineSlice = append(lineSlice, v)
					}
				}
			}

		}

		if lineSlice != nil && len(lineSlice) > 0 {
			for _, lineMap := range lineSlice {
				//lineData := make([]string, len(sheetTitleKeyArr)+baseColLen)
				lineData := baseColVal
				//if baseColLen > 0 {
				//	lineData = append(lineData, baseColVal...)
				//}
				for _, v := range sheetTitleKeyArr {
					//字段在配置文件中指定，是否存在该字段
					value, exist := lineMap[v]
					//第0为保留字段，给非接口字段使用
					if exist {
						lineData = append(lineData, value.(string))
						//lineData[k+baseColLen] = value.(string)
					} else {
						lineData = append(lineData, "")
						//lineData[k+baseColLen] = ""
					}
				}
				sheetSlice = append(sheetSlice, lineData)
			}
		}
	}

	return
}
