package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrorCodeOK        = 0
	ErrorCodeUserCheck = 10000 + iota
	ErrorCodeJsonData
	ErrorCodeAuthInfo
)

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Username string `form:"username" json:"username" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

type AuthInfo struct {
	Token string `form:"token" json:"token"`
}

// TODO 全局变量是否可用 以及命名规则 规范:可变全局变量尽量不用
var users = map[string]string{"francis": "francis", "test": "test"}

func checkUser(name string, password string) bool {
	if user, ok := users[name]; ok {
		if user == password {
			return true
		}
	}
	return false
}

// TODO 需要类似do while 效果
// TODO 出错需要返回HTTP错误码 各种错误如何对应？
func loginHandle(c *gin.Context) {
	var json Login
	var code int
	respSend := make(map[string]interface{})
	for {
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			code = http.StatusBadRequest
			respSend["code"] = ErrorCodeJsonData
			respSend["message"] = err.Error()
			break
		}
		// 判断用户名密码是否正确
		if !checkUser(json.Username, json.Password) {
			code = http.StatusBadRequest
			respSend["code"] = ErrorCodeUserCheck
			respSend["message"] = "user check error"
			break
		}
		// 获取JWT-TOKEN
		jwt := fmt.Sprintf("http://192.168.147.132:9080/apisix/plugin/jwt/sign?key=%v", json.Username)
		if resp, err := http.Get(jwt); nil != err {
			code = http.StatusNotFound
			respSend["code"] = ErrorCodeAuthInfo
			respSend["message"] = err.Error()
		} else {
			defer resp.Body.Close()
			if body, err := io.ReadAll(resp.Body); nil == err {
				code = http.StatusOK
				respSend["token"] = string(body)
			} else {
				code = http.StatusBadRequest
				respSend["code"] = ErrorCodeAuthInfo
				respSend["message"] = err.Error()
			}
		}

		break
	}
	fmt.Println(respSend)
	c.JSON(code, respSend)
}

func loadRoute(e *gin.Engine) {
	e.StaticFile("/", "web/index.html")
	e.StaticFile("/updownload.html", "web/updownload.html")
	e.Static("/script", "web/script")
	e.POST("/login", loginHandle)
}

func routeInitBlock() {
	engine := gin.Default()
	loadRoute(engine)
	engine.Run(":8000")
}
func main() {
	routeInitBlock()
	fmt.Println("------end-------")
}
