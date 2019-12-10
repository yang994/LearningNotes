package main

import (
	"flag"
	"fmt"
	"runtime"
)

type UserState int32

const (
	Starting UserState = iota
	GroupStarted
	BothStarted
)

func main() {
	//a := flag.Int("123", 1, "123?")
	flag.Parse()
	fmt.Println(runtime.GOOS)
}
