package handler

import "testing"

func TestApiMultiHandler(t *testing.T) {
	//url := "http://139.198.23.208:80/rsj/ent/keyword/query_list"
	//apiParams := []map[string]string{
	//	{"key_id": "03136445277747c1","sign_key": "2a6c2430912644fc8258c7be84d3a1ce","keyword": "阿里巴巴", "page_no": "1"},
	//	{"key_id": "03136445277747c1","sign_key": "2a6c2430912644fc8258c7be84d3a1ce","keyword": "腾讯", "page_no": "1"},
	//	{"key_id": "03136445277747c1","sign_key": "2a6c2430912644fc8258c7be84d3a1ce","keyword": "百度", "page_no": "1"},
	//	{"key_id": "03136445277747c1","sign_key": "2a6c2430912644fc8258c7be84d3a1ce","keyword": "字节跳动", "page_no": "1"},
	//}
	//apiMultiHandler(url, apiParams)

	url := "https://rsj.ronglianyiyun.com/rsj/ent/new_reg_info"
	apiParams := []map[string]string{
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-09-05", "local_id": "110000", "page_id": ""},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-09-01", "local_id": "110000", "page_id": ""},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-09-02", "local_id": "110000", "page_id": ""},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-09-03", "local_id": "110000", "page_id": ""},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-10", "end_date": "2020-09-01", "local_id": "110000", "page_id": ""},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-11", "end_date": "2020-09-01", "local_id": "110000", "page_id": ""},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-12", "end_date": "2020-09-01", "local_id": "110000", "page_id": ""},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-13", "end_date": "2020-09-01", "local_id": "110000", "page_id": ""},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-08-15", "local_id": "110000", "page_id": ""},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-08-15", "local_id": "110000", "page_id": ""},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-08-15", "local_id": "110000", "page_id": ""},
		{"key_id": "d042cdb5e6344a0d", "sign_key": "b8945357b8ff470bae6e5a804c2722c1", "start_date": "2020-08-20", "end_date": "2020-08-15", "local_id": "110000", "page_id": ""},
	}
	apiMultiHandler(url, apiParams)

}
