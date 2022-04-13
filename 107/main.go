package main

import (
	"fmt"
	"runtime"
	"runtimetest/sub"
)

func FuncInfo() {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		f := runtime.FuncForPC(pc)
		fmt.Printf("file:%s, line:%d,func:%s\n", file, line, f.Name())
		fmt.Printf("%s:%d,func:%s\n", file, line, f.Name())
	} else {
		fmt.Printf("Caller not ok")
	}

}

func Callers() {
	sub.Sub()
}

func main() {
	FuncInfo()
	fmt.Println("----------------------")
	Callers()
}
