package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// HTTP get
func HttpGet(url string) ([]byte, error) {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("HttpGet error:%s\n", err.Error())
		return nil, err
	}
	fmt.Printf("status: %d\n", resp.StatusCode)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Printf("response body:%s\n", string(body))

	if resp.StatusCode == http.StatusOK {
		return body, nil
	} else {
		fmt.Printf("HttpGet error status:%d, body:%s\n", resp.StatusCode, string(body))
		return nil, errors.New(string(body))
	}
}

// http post
func HttpPostJSON(url string, data interface{}) ([]byte, error) {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Marshal error %s\n", err.Error())
		return nil, err
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Printf("HttpPostJSON error:%s\n", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		fmt.Println("ok")
		return body, nil
	} else {
		fmt.Printf("HttpPostJSON error status:%d, body:%s\n", resp.StatusCode, string(body))
		return nil, errors.New(string(body))
	}
}

// http delete
func HttpDelete(url string) ([]byte, error) {
	fmt.Printf("HttpDelete %s\n", url)
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}
	fmt.Printf("status: %d\n", resp.StatusCode)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		return body, nil
	} else {
		fmt.Printf("HttpDelete error status:%d, body:%s\n", resp.StatusCode, string(body))
		return nil, errors.New(string(body))
	}
}

// http put
func HttpPutJSON(url string, data interface{}) ([]byte, error) {
	fmt.Printf("HttpPutJSON %s\n", url)

	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}
	client := &http.Client{Timeout: 5 * time.Second}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}
	fmt.Printf("status: %d\n", resp.StatusCode)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		return body, nil
	} else {
		fmt.Printf("HttpPutJSON error status:%d, body:%s\n", resp.StatusCode, string(body))
		return nil, errors.New(string(body))
	}
}

func loop1000(host string, port int, loop int) {
	start := time.Now()
	req := fmt.Sprintf("http://%v:%v/aa/bb", host, port)
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
	host := "127.0.0.1"
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

/*******************************不是很好的写法示例******************************************/

// 不是很好的 发送GET请求示例
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
	body := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		body.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return body.String()
}

// 不是很好的 发送POST请求示例
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
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
