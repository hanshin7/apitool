package utils

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	//println(fmt.Sprint(time.Now().Unix()))
	println(time.Now().Format("20060102"))
	stampStr := fmt.Sprint(time.Now().Unix())
	println(stampStr)
	println("TEST" + time.Now().Format("20060102") + stampStr + fmt.Sprint(rand.Intn(1000)))

}

func Test2(t *testing.T) {
	allResp := []map[string]interface{}{}
	a := apier{url: "/123", allResp: &allResp}
	println(len(allResp))
	a.apitest()
	println(len(allResp))

}

type apier struct {
	url     string
	allResp *[]map[string]interface{}
}

func (api *apier) apitest() {
	m := make(map[string]interface{})
	m["k"] = "v"
	*api.allResp = append(*api.allResp, m)
}
