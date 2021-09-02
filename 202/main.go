package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.StaticFile("/", "web/index.html")
	router.StaticFile("/favicon.ico", "web/favicon.ico")
	router.Static("/static", "resource")
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		log.Println(form)
		files := form.File["files"] // the input tag name attribute

		for _, file := range files {
			log.Println(file.Filename)
			log.Println(file.Header)
			log.Println(file.Size)

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, "resource"+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	router.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf(`{code:0, message:"OK"}`))
	})
	router.Run(":8082")
}
