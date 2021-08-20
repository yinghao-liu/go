package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	ErrorCodeOK        = 0
	ErrorCodeUserCheck = 10000 + iota
	ErrorCodeJsonData
	ErrorCodeAuthInfo
	ErrorCodeFiles
)

type FileInfo struct {
	Name    string
	Size    int64
	ModTime string
}

func filesHandle(c *gin.Context) {
	respSend := make(map[string]interface{})
	if a, err := os.ReadDir("resource"); nil == err {
		filesInfo := make([]FileInfo, len(a))
		for i, v := range a {
			fmt.Println(v.Name())
			filesInfo[i].Name = v.Name()
			if info, err := v.Info(); nil == err {
				fmt.Printf("%v-%v-%v\n", info.Name(), info.Size(), info.ModTime())
				filesInfo[i].Size = info.Size()
				filesInfo[i].ModTime = info.ModTime().String()
			}
		}
		fmt.Println(filesInfo)
		respSend["code"] = ErrorCodeOK
		respSend["message"] = ""
		respSend["data"] = filesInfo
	} else {
		respSend["code"] = ErrorCodeFiles
		respSend["message"] = err.Error()
	}
	c.JSON(http.StatusOK, respSend)
}

func fileGetHandle(c *gin.Context) {
	name := c.Param("filename")
	fullpath := "resource/" + name
	fmt.Printf("name is %v\n", fullpath)
	c.Header("Content-Type", "application/octet-stream")
	c.File(fullpath)
}

func filePostHandle(c *gin.Context) {
	respSend := make(map[string]interface{})
	if a, err := os.ReadDir("resource"); nil == err {
		number := len(a)
		filesInfo := make([]FileInfo, number)

		for i, v := range a {
			fmt.Println(v.Name())
			filesInfo[i].Name = v.Name()
			if info, err := v.Info(); nil == err {
				fmt.Printf("%v-%v-%v\n", info.Name(), info.Size(), info.ModTime())
				filesInfo[i].Size = info.Size()
				filesInfo[i].ModTime = info.ModTime().String()
			}
		}
		fmt.Println(filesInfo)
		respSend["code"] = ErrorCodeOK
		respSend["message"] = ""
		respSend["data"] = filesInfo
	} else {
		respSend["code"] = ErrorCodeFiles
		respSend["message"] = err.Error()
	}
	c.JSON(http.StatusOK, respSend)
}

func fileDeleteHandle(c *gin.Context) {
	respSend := make(map[string]interface{})
	name := c.Param("filename")
	fullpath := "resource/" + name
	fmt.Printf("name is %v\n", fullpath)
	respSend["code"] = ErrorCodeOK
	respSend["message"] = ""
	if err := os.Remove(fullpath); nil != err {
		respSend["code"] = ErrorCodeFiles
		respSend["message"] = err.Error()
	}
	c.JSON(http.StatusOK, respSend)
}

func loadRoute(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		v1.GET("/files", filesHandle)
		v1.GET("/files/:filename", fileGetHandle)
		v1.POST("/files/:filename", filePostHandle)
		v1.DELETE("/files/:filename", fileDeleteHandle)
	}
}

func routeInitBlock() {
	engine := gin.Default()
	loadRoute(engine)
	engine.Run(":8002")
}

func main() {
	routeInitBlock()
	fmt.Println("------end-------")
}
