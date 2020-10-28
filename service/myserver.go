package service

import (
	"apitool/config"
	"apitool/handler"
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//启动服务
func StartService() {
	serverPort := config.Conf.Section("sys").Key("server_port").Value()
	//开放file目录访问权限,可供页面下载
	fs := http.FileServer(http.Dir("static/file"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/fileQuery", uploadHandler)
	fmt.Printf("服务启动，监听端口为[%s]\n", serverPort)
	http.ListenAndServe("localhost:"+serverPort, nil)
}

//停止服务
func shutdownService() {

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("page/index.tpl")
	t.Execute(w, nil)
}

/**
* 批量文件查询接口处理
 */
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 >> 20)
	file, h, err := r.FormFile("ufile")
	if err != nil {
		errors.New("文件上传错误!")
		return
	}
	path := config.Conf.Section("sys").Key("file_path").Value()
	filePath := path + "/upload/" + h.Filename
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		errors.New("文件上传保存错误!")
		return
	}
	io.Copy(f, file)
	f.Close()
	file.Close()
	fmt.Printf("表单上传文件[%s]已保存\n", h.Filename)

	//解析表单域
	r.ParseForm()
	apipath := r.Form["apipath"][0]
	//执行接口调用前数据解析处理
	apiParams, err := parseUploadFile(filePath, r.Form)
	if err != nil {
		return
	}
	//执行接口调用逻辑
	fileName, err := handler.ApiMultiHandler(apipath, apiParams)
	if err != nil {
		return
	}
	fmt.Printf("批量查询完成，结果文件名为：%s\n", fileName)

	t, _ := template.ParseFiles("page/entfile_result.tpl")
	t.Execute(w, "success")
}

/**
* 解析上传的接口请求参数文件，保存到map
 */
func parseUploadFile(filePath string, params url.Values) (apiParams []map[string]string, err error) {
	pmykey := params["pmykey"][0]
	psignkey := params["psignkey"][0]

	//解析文件到map
	apiParams = []map[string]string{}
	f, err := os.Open(filePath)
	if err != nil {
		err = errors.New("打开文件失败:" + err.Error())
		return
	}
	r := bufio.NewReader(f)
	//读有第一行表头数据，取key
	headLine, _, err := r.ReadLine()
	if err != nil {
		err = errors.New("打开文件失败:" + err.Error())
		return
	}
	headLineStr := strings.TrimSpace(string(headLine))
	headLineArr := strings.Split(headLineStr, ",")
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			//err = errors.New("解析文件失败:" + err.Error())
			//return
		}
		lineStr := strings.TrimSpace(string(line))
		lineArr := strings.Split(lineStr, ",")
		params := make(map[string]string)
		//添加api业务字段
		for i, v := range lineArr {
			params[headLineArr[i]] = v
		}
		//密钥字段
		params["key_id"] = pmykey
		params["sign_key"] = psignkey
		apiParams = append(apiParams, params)
	}
	return
}
