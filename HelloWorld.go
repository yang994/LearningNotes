package main

import (
	"fmt"
	"strconv"

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

func main() {
	a := i1{
		name: "123",
	}
	a.testfunc2("345")
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
