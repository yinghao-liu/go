package main

import (
	"fmt"
	"net/url"
)

func urlGenerate() {
	var a url.URL
	a.Scheme = "http"
	a.Host = "127.0.0.1:8000"
	//a.Path = "/v1/account?a=b"
	//a.RawPath = url.PathEscape("/v1/account?a=b")

	q := a.Query()
	q.Add("aaa", "b#bb")
	a.RawQuery = q.Encode()
	fmt.Printf("---url is %s\n", a.String())
	raw, err := url.PathUnescape(a.String())
	if nil != err {
		fmt.Printf("err is %s\n", err.Error())
	}
	fmt.Printf("---url raw is %s\n", raw)

}

func rtspParser() {
	u, err := url.Parse("rtsp://192.168.0.4:554/realtime?a=b")
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("u is %+v\n", *u)
	fmt.Printf("u string is %s\n", u.String())
	i := url.UserPassword("admin", "francis")
	if nil != i {
		fmt.Printf("i is %+v\n", *i)
		fmt.Printf("i string is %s\n", i.String())
	}
	u.User = i
	fmt.Printf("u after user string is %s\n", u.String())
}

func main() {
	uoo := url.PathEscape("/v1/account?a=b")
	fmt.Println(uoo)

	url2, ok := url.PathUnescape("/%3CplatformUrl%3E/allNotify/")
	fmt.Println(url2, ok)

	urlGenerate()

	rtspParser()
}
