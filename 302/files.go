package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// 错误码，需要统一定义
const (
	ErrorCodeOK        = 0            // 成功
	ErrorCodeUserCheck = 10000 + iota //用户名密码检查
	ErrorCodeJsonData                 // json 数据错误
	ErrorCodeAuthInfo                 // 认证信息错误
	ErrorCodeFiles                    //文件操作错误
)

// 文件信息
type FileInfo struct {
	Name    string `json:"name"`    // 文件名
	Size    int64  `json:"size"`    // 文件大小（字节）
	ModTime string `json:"modTime"` // 修改时间
}

// 线程安全的map
var fileRWLock sync.Map

// 获取文件列表接口
func filesHandle(c *gin.Context) {
	var code int = http.StatusOK
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
		respSend["fileList"] = filesInfo
	} else {
		code = http.StatusBadRequest
		respSend["code"] = ErrorCodeFiles
		respSend["message"] = err.Error()
	}
	c.JSON(code, respSend)
}
func listSyncMap(c *gin.Context) {
	fmt.Println("------listSyncMap---------")
	fileRWLock.Range(func(key, value interface{}) bool {
		fmt.Printf("%v, %v\n", key, value)
		return true
	})
	c.Status(http.StatusOK)
	fmt.Println("------listSyncMap end---------")
}

// 获取文件内容接口
func fileGetHandle(c *gin.Context) {
	name := c.Param("filename")
	fullpath := "resource/" + name
	fmt.Printf("name is %v\n", fullpath)

	c.Header("Content-Type", "application/octet-stream")
	//c.Header("Content-Disposition", "attachment; filename="+name)

	value, _ := fileRWLock.LoadOrStore(fullpath, new(sync.RWMutex))
	var rwLock *sync.RWMutex = value.(*sync.RWMutex)
	rwLock.RLock() // 写锁
	defer rwLock.RUnlock()

	c.File(fullpath)
}

// 创建文件接口
func filePostHandle(c *gin.Context) {
	// form, _ := c.MultipartForm()
	// files := form.File["upload"]

	// for _, file := range files {
	// 	log.Println(file.Filename)

	// 	// Upload the file to specific dst.
	// 	c.SaveUploadedFile(file, dst)
	// }
	// c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	name := c.Param("filename")
	var code int = http.StatusOK
	var contentLen int
	var err error
	fullpath := "resource/" + name
	fmt.Printf("name is %v\n", fullpath)

	respSend := make(map[string]interface{})
	contentLenStr := c.GetHeader("Content-Length")
	fmt.Printf("Content-Length is %v\n", contentLenStr)

	if contentLen, err = strconv.Atoi(contentLenStr); nil == err {
		fmt.Printf("contentLen is %v\n", contentLen)
	}
	// 获取http传输过来的完整内容
	var content = make([]byte, contentLen)
	if content, err = c.GetRawData(); nil != err {
		fmt.Println(err)
		code = http.StatusBadRequest
		respSend["code"] = ErrorCodeFiles
		respSend["message"] = err.Error()
		return
	}

	value, _ := fileRWLock.LoadOrStore(fullpath, new(sync.RWMutex))
	var rwLock *sync.RWMutex = value.(*sync.RWMutex)
	rwLock.Lock() // 写锁
	defer rwLock.Unlock()

	// 获取内容写到文件中
	// TODO 添加文件是否存在的判断
	if file, err := os.OpenFile(fullpath, os.O_RDWR|os.O_CREATE, 0755); err != nil {
		log.Fatal(err)
		code = http.StatusBadRequest
		respSend["code"] = ErrorCodeFiles
		respSend["message"] = err.Error()
	} else {
		defer file.Close()
		if writeLen, err := file.Write(content); nil != err {
			code = http.StatusBadRequest
			respSend["code"] = ErrorCodeFiles
			respSend["message"] = err.Error()
		} else {
			fmt.Printf("writeLen: %v\n", writeLen)
			fmt.Printf("contentLen: %v\n", contentLen)
			if writeLen != contentLen {
				code = http.StatusBadRequest
				respSend["code"] = ErrorCodeFiles
				respSend["message"] = fmt.Sprintf("writeLen not match wirten:%v, source:%v", writeLen, contentLen)
			} else {
				code = http.StatusCreated
			}
		}
	}

	if http.StatusCreated != code {
		c.JSON(code, respSend)
	} else {
		c.Status(code)
	}
}

// 文件删除操作
func fileDeleteHandle(c *gin.Context) {
	var code int = http.StatusOK
	respSend := make(map[string]interface{})
	name := c.Param("filename")
	fullpath := "resource/" + name
	fmt.Printf("name is %v\n", fullpath)

	value, _ := fileRWLock.LoadOrStore(fullpath, new(sync.RWMutex))
	var rwLock *sync.RWMutex = value.(*sync.RWMutex)
	rwLock.Lock() // 写锁
	defer rwLock.Unlock()
	defer fileRWLock.Delete(fullpath)

	if err := os.Remove(fullpath); nil != err {
		code = http.StatusBadRequest
		respSend["code"] = ErrorCodeFiles
		respSend["message"] = err.Error()
	}

	if http.StatusOK != code {
		c.JSON(code, respSend)
	} else {
		c.Status(code)
	}
}

// 路由定义
func loadRoute(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		v1.GET("/files", filesHandle)
		v1.GET("/files/:filename", fileGetHandle)
		v1.POST("/files/:filename", filePostHandle)
		v1.DELETE("/files/:filename", fileDeleteHandle)
	}
	e.GET("/syncMap", listSyncMap)
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
