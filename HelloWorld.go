package main

import (
	"fmt"

	shell "github.com/xcshuan/go-mefs-api"
)

func main() {
	a := 123123 * 1024 * 1024
	fmt.Println(a / 1024 / 1024 / 60 / 60 * 10)
}

func testRS() {
	shell.ReslultSumaryTest("33373", "41667")
}
