package main

import (
	"fmt"
	"net/url"
)

func urlGenerate() {
	var a url.URL
	a.Scheme = "http"
	a.Host = "127.0.0.1:8000"
	a.Path = "/v1/account"
	fmt.Printf("url is %s\n", a.String())

}

func main() {
	uoo := url.PathEscape("<dasdsd>/sdfsfdsdfs")
	fmt.Println(uoo)

	url2, ok := url.PathUnescape("/%3CplatformUrl%3E/allNotify/")
	fmt.Println(url2, ok)

	urlGenerate()
}
