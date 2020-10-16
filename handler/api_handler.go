package handler

import (
	"apitool/utils"
	"apitool/workpool"
	"fmt"
	"sync"
)

//并发锁
var mutex sync.Mutex

/**
* 接口并发请发处理,并发量通过线程池大小控制
* url:请求api完整路径
* apiParams：多条请求参数，切片存储，每一条切片值对应一次api参数
 */
func apiMultiHandler(url string, apiParams []map[string]string) {
	//保存所有请求的响应结果,多线程些，需处理并发安全
	allResp := []string{}
	// 使用5个 goroutine 来创建工作池
	p := workpool.New(2)
	var wg sync.WaitGroup
	//数量为需调用接口条数
	wg.Add(len(apiParams))
	fmt.Printf("workpool=%d tasknum=%d\n", 2, len(apiParams))
	//循环并发调用接口
	for _, apiParamMap := range apiParams {
		//创建api调用类型体
		api := apiSender{
			url:         url,
			apiParamMap: apiParamMap,
			allResp:     &allResp,
		}

		//接口并发调用，最大并发个数由工作池限制
		go func() {
			p.Run(&api)
			wg.Done()
		}()
	}
	wg.Wait()
	p.Shutdown()

	//加工excel数据
	exData := dealResp2ExcelData(url, &allResp)
	println("datalen:", len(exData.SheetDataSlice[0]))
	//生成excel
	filename, err := utils.CreateExcel(exData)
	if err != nil {

	} else {
		println(filename)
	}

}

//接口响应数据转化成excel数据
func dealResp2ExcelData(url string, resps *[]string) *utils.ExcelData {
	fmt.Printf("api result num：%d\n", len(*resps))
	//创建Excel数据类型结构体
	exData := utils.ExcelData{
		ExcelName:       utils.Url2Name(url),
		SheetNameSlice:  []string{"接口响应报文"},
		SheetTitleSlice: [][]string{{"请求参数", "响应结果"}},
	}

	//接口响应报文sheet数据
	sheetSlice := [][]string{}
	//接口响应数据转化成excel数据
	for _, resp := range *resps {
		lineSlice := []string{resp}
		//本条数据添加到对应sheet
		sheetSlice = append(sheetSlice, lineSlice)
	}
	exData.SheetDataSlice = [][][]string{sheetSlice}
	return &exData
}

type apiSender struct {
	url         string
	apiParamMap map[string]string
	allResp     *[]string
}

// Task 实现 Worker 接口
func (api *apiSender) Task() {
	//处理接口调用业务逻辑
	resp := utils.RequestApi(api.url, api.apiParamMap)
	//将每个请求返回结果保存到切片
	mutex.Lock()
	//指针，多个线程向主线程同一个切片写值
	*api.allResp = append(*api.allResp, resp)
	mutex.Unlock()
}
