package model

//接口响应报文结构体
type ApiRespMsg struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageResult struct {
	Code        string
	Msg         string
	CsvFileName string
	XlsFileName string
}

type SingleQueryResult struct {
	Code string
	Msg  string
	Data string
}
