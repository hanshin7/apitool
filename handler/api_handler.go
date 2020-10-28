package handler

import (
	"apitool/config"
	"apitool/utils"
	"apitool/workpool"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

//并发锁
var mutex sync.Mutex

/**
* 接口并发请发处理,并发量通过线程池大小控制
* url:请求api路径,不包含ip和端口地址
* apiParams：多条请求参数，切片存储，每一条切片值对应一次api参数
 */
func ApiMultiHandler(url string, apiParams []map[string]string) (fileName string, err error) {
	//保存所有请求的响应结果,多线程些，需处理并发安全
	allResp := []string{}
	num, err := strconv.Atoi(config.Conf.Section("sys").Key("workpool_num").Value())
	if err != nil {
		return
	}
	// 使用n个 goroutine 来创建工作池
	p := workpool.New(num)
	var wg sync.WaitGroup
	//数量为需调用接口条数
	wg.Add(len(apiParams))
	fmt.Printf("workpool=%d tasknum=%d\n", num, len(apiParams))
	url_path := config.Conf.Section("sys").Key("http_path").Value() + url
	//循环并发调用接口
	for _, apiParamMap := range apiParams {
		//创建api调用类型体
		api := apiSender{
			url:         url_path,
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

	fileName = config.Conf.Section("url").Key(url).Value() + fmt.Sprint(time.Now().Unix())
	//保存接口响应报文到文件
	wfErr := utils.WriteFile(config.Conf.Section("sys").Key("file_path").Value()+"/download/"+fileName+".csv", &allResp)
	if wfErr != nil {
		return
	}

	//解析接口数据生成excel文件数据格式,根据不同接口定义不同解析规则
	exData := utils.ParseRespData(url, fileName, &allResp)
	//生成excel
	ok := utils.CreateExcel(exData)
	if !ok {
		err = errors.New("生成excel文件失败")
		return
	}

	return
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
