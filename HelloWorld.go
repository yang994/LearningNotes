package main

import (
	"fmt"
	"sync"

	shell "github.com/ipfs/go-mefs-api"
)

func main() {
	var test sync.Map
	test.Store("aaa", "bbb")
	testvalue, ok := test.Load("aaa")
	if ok {
		fmt.Println("testvalue:", testvalue)
	}
	test.Delete("a")
	testvalue, ok = test.Load("aaa")
	if ok {
		fmt.Println("testvalue:", testvalue)
	} else {
		fmt.Println("删除成功,testvalue:", testvalue)
	}

}

func testRS() {
	shell.ReslultSumaryTest("33373", "41667")
}
