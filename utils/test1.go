package utils

import "fmt"

func init() {
	fmt.Println("init function --->")
}

var _int64 = s()

func s() int64 {
	fmt.Println("function s() --->")
	return 1
}

func main() {
	fmt.Println("main --->")
}
