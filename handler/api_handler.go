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
	//allResp := [][]map[string]interface{}{}
	// 使用5个 goroutine 来创建工作池
	p := workpool.New(2)
	var wg sync.WaitGroup
	//数量为需调用接口条数
	wg.Add(len(apiParams))
	fmt.Printf("workpool=%d tasknum=%d\n", 2, len(apiParams))
	//循环并发调用接口
	for _, apiParamMap := range apiParams {
		//创建api调用类型体
		api := apier{
			url:         url,
			apiParamMap: apiParamMap,
			//allResp: &allResp,
		}

		//接口并发调用，最大并发个数由工作池限制
		go func() {
			p.Run(&api)
			wg.Done()
		}()
	}
	wg.Wait()
	p.Shutdown()

}

type apier struct {
	url         string
	apiParamMap map[string]string
	//allResp *[][]map[string]interface{}
}

// Task 实现 Worker 接口
func (api *apier) Task() {
	//处理接口调用业务逻辑
	resp := utils.RequestApi(api.url, api.apiParamMap)
	//处理接口返回值
	mutex.Lock()
	//api.allResp = append(api.allResp, resp)
	mutex.Unlock()

}
