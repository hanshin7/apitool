package service

import (
	"html/template"
	"net/http"
)

//启动服务
func StartService() {
	http.HandleFunc("/page/index", indexHandler)
	http.ListenAndServe("localhost:8089", nil)

}

//停止服务
func shutdownService() {

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("page/index.tpl")
	t.Execute(w, nil)
}
