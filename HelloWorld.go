package main

import (
	"fmt"
	"strconv"
	"strings"

	shell "github.com/xcshuan/go-mefs-api"
)

func main() {
	a := strings.Split("//123/45", "/")
	for _, b := range a {
		fmt.Println(b)
	}
}

func testRS() {
	shell.ReslultSumaryTest("33373", "41667")
}

func testfunc(str string, args ...interface{}) {
	a := str
	for _, operator := range args {
		switch operator.(type) {
		case string:
			a += "(string)" + operator.(string)
		case int:
			a += "(int)" + strconv.Itoa(operator.(int))
		default:
		}
	}
	fmt.Println(a)
}
