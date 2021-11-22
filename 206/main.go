package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// HTTP代理
func HTTPProxy(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("id is %s\n", id)
	var ext string
	if 0 != len(id) {
		ext = fmt.Sprintf("/%s", id)
	}

	host := "127.0.0.1:8000"
	resUrl := "/v"
	var u url.URL
	u.Scheme = "http"
	u.Host = host
	fmt.Printf("%+v\n", u)
	fmt.Printf("%+s\n", u.String())
	proxy := httputil.NewSingleHostReverseProxy(&u)
	proxy.Director = func(req *http.Request) {
		req.URL.Path = resUrl + ext
		req.URL.Scheme = "http"
		req.URL.Host = host
		req.Host = host
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	r := gin.Default()

	catalogs := r.Group("/v")
	{
		catalogs.GET("/:id", HTTPProxy)
	}
	// 监听端口，默认在8080
	r.Run(":8000")
}
