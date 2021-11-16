package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type VideoType struct {
	Type string `json:"type" binding:"required,oneof=homemade reprinted"`
}
type Videos struct {
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	Category     string `json:"category"`
	VideoType
	Config *interface{} `json:"config"`
}

type Homemade struct {
	Permission string `json:"permission"`
}
type Reprinted struct {
	Src string `json:"src"`
}

// 视频添加
func ProductAdd(c *gin.Context) {
	var t VideoType
	if err := c.ShouldBindBodyWith(&t, binding.JSON); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "message": err.Error()})
		return
	}
	fmt.Printf("t is %+v\n", t)

	var v Videos
	switch t.Type {
	case "homemade":
		var h Homemade
		v.Config = h
	case "reprinted":
		var r Reprinted
		v.Config = r
	}

	var h Homemade
	v.Config = &h
	if err := c.ShouldBindBodyWith(&v, binding.JSON); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": 1001, "message": err.Error()})
		return
	}
	fmt.Printf("v is %+v\n", v)

	c.JSON(http.StatusOK, "{}")

}

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.POST("/videos", ProductAdd)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
