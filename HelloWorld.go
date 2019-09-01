package main

import (
	"fmt"
	"runtime"
	"strings"

	shell "github.com/xcshuan/go-mefs-api"
)

type UserState int32

const (
	Starting UserState = iota
	GroupStarted
	BothStarted
)

func main() {
	fmt.Println(strings.HasPrefix("abc", ""))
}

func getCaller() {
	for i := range []int{0, 1, 2, 3, 4} {
		pc, _, _, _ := runtime.Caller(i)
		fmt.Println(i, runtime.FuncForPC(pc).Name())
	}
}

func getprice(spacetime int, amount int64) int64 {
	return amount * 24 * 60 * 60 * 1024 * 1024 / int64(spacetime)
}

func convertSpacetime(spacetime int, price int64) int64 {
	if spacetime <= 0 {
		fmt.Println("error! spaceTime:", spacetime)
		return 0
	}
	amount := int64(spacetime) * price / 1024 / 1024 / 60 / 60 / 24 //注意这里先用时空值×单位，计算出来更加准确
	if amount <= 0 {
		fmt.Println("error! spaceTime:", spacetime, "amount:", amount)
		return 0
	}
	return amount
}

func testRS() {
	shell.ReslultSumaryTest("33373", "41667")
}
