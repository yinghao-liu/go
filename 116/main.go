package main

import "fmt"

func main() {
	killProcessByName("abc")
	if err := startProcessIfNotExist("abc", nil); nil != err {
		fmt.Printf("startProcessIfNotExist error is %s\n", err.Error())
	}

}
