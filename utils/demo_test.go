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
