package main

import (
	"fmt"
	"net/url"
)

func main() {
	uoo := url.PathEscape("<dasdsd>/sdfsfdsdfs")
	fmt.Println(uoo)

	url2, ok := url.PathUnescape("/%3CplatformUrl%3E/allNotify/")
	fmt.Println(url2, ok)
}
