package handler

import (
	"apitool/rules"
	"apitool/utils"
	"testing"
)

func TestApiMultiHandler(t *testing.T) {
	//url := "http://139.198.23.208:80/rsj/ent/keyword/query_list"
	url := "rsj/ent/keyword/query_list"
	apiParams := []map[string]string{
		{"key_id": "03136445277747c1", "sign_key": "2a6c2430912644fc8258c7be84d3a1ce", "keyword": "阿里巴巴", "page_no": "1"},
		{"key_id": "03136445277747c1", "sign_key": "2a6c2430912644fc8258c7be84d3a1ce", "keyword": "腾讯", "page_no": "1"},
		//{"key_id": "03136445277747c1","sign_key": "2a6c2430912644fc8258c7be84d3a1ce","keyword": "百度", "page_no": "1"},
		//{"key_id": "03136445277747c1","sign_key": "2a6c2430912644fc8258c7be84d3a1ce","keyword": "字节跳动", "page_no": "1"},
	}
	ApiMultiHandler(url, apiParams)

	//url := "https://rsj.ronglianyiyun.com/rsj/ent/new_reg_info"
	//apiParams := []map[string]string{
	//	{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-09-05", "local_id": "110000", "page_id": ""},
	//	{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-09-01", "local_id": "110000", "page_id": ""},
	//	{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-09-02", "local_id": "110000", "page_id": ""},
	//}
	//apiMultiHandler(url, apiParams)

}

func TestApiMultiHandler2(t *testing.T) {
	url := "rsj/person/passengerstatid/query"
	apiParams := []map[string]string{
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "name": "韩威", "pid": "420117199408201239", "gid": "", "i_month": "12"},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "name": "韩威", "pid": "420117199408201239", "gid": "", "i_month": "12"},
	}
	ApiMultiHandler(url, apiParams)

	//url := "https://rsj.ronglianyiyun.com/rsj/ent/new_reg_info"
	//apiParams := []map[string]string{
	//	{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-09-05", "local_id": "110000", "page_id": ""},
	//	{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-09-01", "local_id": "110000", "page_id": ""},
	//	{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-09-02", "local_id": "110000", "page_id": ""},
	//}
	//apiMultiHandler(url, apiParams)

}

func TestParseGRGS(t *testing.T) {
	//config.InitConfig("../config/config.ini")
	////初始化日志配置
	//level := config.Conf.Section("sys").Key("log_level").Value()
	//logging.Init(level)
	apiParams := []map[string]string{}
	allResp := []string{"{\"cert\":\"e263405854ff684969bef1a81990aed3\",\"code\":\"RSJ000\",\"data\":{\"ryposshas\":[],\"ryposfrs\":[{\"esdate\":\"2012-10-18\",\"regno\":\"440101400124320\",\"creditcode\":\"91440101054546260E\",\"entname\":\"南洋商业银行(中国)有限公司广州越秀支行\",\"enttype\":\"台、港、澳投资企业分公司\",\"regcap\":\"\",\"palgorithmid\":\"E263405854FF684969BEF1A81990AED3\",\"ryname\":\"冯君\",\"regorgprovince\":\"广东省\",\"industryphyname\":\"金融业\",\"entstatus\":\"注销\",\"regcapcur\":\"人民币元\"}],\"entcaseinfos\":[],\"punisheds\":[],\"punishbreaks\":[],\"rypospers\":[]},\"msg\":\"成功\",\"response_sign\":\"460ccbd29a5803852d90b2ed9fec7430a8b719a084f07065a3160265c0f16db1\"}"}
	exData := rules.NewParseApiStrategy("rsj/person/gsinfo/query", "test111", apiParams, allResp)
	//生成excel
	ok := utils.CreateExcel(exData)
	println(ok)
}

//func TestMain(m *testing.M) {
//	//m.Run()
//	//config.InitConfig("../config/config.ini")
//	//初始化日志配置
//	level := config.Conf.Section("sys").Key("log_level").Value()
//	logging.Init(level)
//}
