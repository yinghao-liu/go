package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func loop1000(host string, port int, loop int) {
	start := time.Now()
	req := fmt.Sprintf("http://%v:%v/decoder/publisher", host, port)
	fmt.Printf("port %v, loop %v, req %v\n", port, loop, req)
	var res string
	for i := 0; i < loop; i++ {
		res = Get(req)
	}
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Printf("total %v ms, %v us\n", elapsed.Milliseconds(), elapsed.Microseconds())
	fmt.Printf("average %v ms, %v us\n", elapsed.Milliseconds()/int64(loop), elapsed.Microseconds()/int64(loop))
}

func main() {
	port := 80
	loop := 1000
	host := "172.16.199.2"
	fmt.Printf("1-- host is %v, port is %v, loop is %v\n", host, port, loop)
	if len(os.Args) > 1 {
		host = os.Args[1]

		i, err := strconv.Atoi(os.Args[2])
		if nil == err {
			fmt.Printf("port is %v\n", i)
			port = i
		}

		i, err = strconv.Atoi(os.Args[3])
		if nil == err {
			fmt.Printf("loop is %v\n", i)
			loop = i
		}
	}
	fmt.Printf("2-- host is %v, port is %v, loop is %v\n", host, port, loop)
	loop1000(host, port, loop)
}
