package main

import (
	"fmt"
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
	Token string
}

// TODO 全局变量是否可用 以及命名规则 规范:可变全局变量尽量不用
var users = map[string]string{"francis": "francis"}

func checkUser(name string, password string) bool {
	if user, ok := users[name]; ok {
		if user == password {
			return true
		}
	}
	return false
}

// TODO 需要类似do while 效果
func loginHandle(c *gin.Context) {
	var json Login
	respSend := make(map[string]interface{})
	for {
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			respSend["code"] = ErrorCodeJsonData
			respSend["message"] = err.Error()
			break
		}
		// 判断用户名密码是否正确
		if !checkUser(json.Username, json.Password) {
			respSend["code"] = ErrorCodeUserCheck
			respSend["message"] = "user check error"
			break
		}

		break
	}
	c.JSON(http.StatusOK, respSend)
}

func loadRoute(e *gin.Engine) {
	e.StaticFile("/", "index.html")
}

func routeInitBlock() {
	engine := gin.Default()
	loadRoute(engine)
	engine.Run("127.0.0.1:80")
}
func main() {
	routeInitBlock()
	fmt.Println("------end-------")
}
