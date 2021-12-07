package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cm Claims
		auth := c.GetHeader("Authorization")
		token, err := ParseToken(auth, &cm)
		if nil == err && token.Valid {
			fmt.Printf("token is %+v\n", token)
			fmt.Printf("Claims is %+v\n", token.Claims)
			fmt.Printf("Claims2 is %+v\n", cm)
		} else {
			fmt.Printf("error is %+v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"code": 1001, "message": err.Error()})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

// 路由定义
func loadRoute(e *gin.Engine) {
	v1file := e.Group("/v1/files")
	v1file.Use(MiddleWare())
	{
		v1file.GET("", filesHandle)
		v1file.GET("/:filename", fileGetHandle)
		v1file.POST("/:filename", filePostHandle)
		v1file.DELETE("/:filename", fileDeleteHandle)
	}

	e.StaticFile("/", "web/index.html")
	e.Static("/web", "web")
	e.POST("/login", loginHandle)

	e.GET("/syncMap", listSyncMap)
}

func routeInitBlock() {
	engine := gin.Default()
	loadRoute(engine)
	engine.Run(":8001")
}
func main() {
	var cm Claims
	token_string, err := GenerateToken("francis")
	if nil == err {
		fmt.Printf("token_string is %s\n", token_string)
		token, err := ParseToken(token_string, &cm)
		if nil == err {
			fmt.Printf("token is %+v\n", token)
			fmt.Printf("Claims is %+v\n", token.Claims)
			fmt.Printf("Claims2 is %+v\n", cm)
		} else {
			fmt.Printf("error is %+v\n", err)
		}
	}
	routeInitBlock()
	fmt.Println("------end-------")
}
