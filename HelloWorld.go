package main

import (
	"fmt"
	"strconv"
	"time"

	shell "github.com/xcshuan/go-mefs-api"
)

type testI interface {
	testfunc1()
	testfunc2() string
}

type i1 struct {
	name string
}

func (p *i1) testfunc1() {
	fmt.Println("i1 name:", p.name)
}

func (p *i1) testfunc2(rename string) string {
	p.name = rename
	fmt.Println("i1 name:", p.name)
	return rename
}

type UserState int32

const (
	Starting UserState = iota
	GroupStarted
	BothStarted
)

func main() {
	testGoRouting()
	fmt.Println("gorouting函数已经返回")
	time.Sleep(time.Hour)
}

func testGoRouting() {
	go func() {
		for i := 1; ; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()
	return
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
