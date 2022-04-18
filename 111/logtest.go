package main

import (
	"fmt"
	"log"
)

func LogTest() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	fmt.Println(log.Flags())
	log.Printf("aaaaaaa")
	log.Output(2, "bbbbbb")
}
func main() {

}
