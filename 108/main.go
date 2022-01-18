package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Validator struct {
	Mac   string `json:"mac" binding:"mac"`
	Email string `json:"email" binding:"email"`
	Ascii string `json:"ascii" binding:"min=6,max=10,printascii"`
}

func servicePutValidator(c *gin.Context) {
	var valid Validator
	if err := c.ShouldBindJSON(&valid); err != nil {
		fmt.Printf("valid: %+v\n", valid)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf("valid: %+v\n", valid)

	c.Status(http.StatusOK)

}
func serviceVersionGet(c *gin.Context) {

	c.Status(http.StatusOK)

}
func main() {

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 服务类
	srv := r.Group("/service")
	{
		srv.PUT("/validator", servicePutValidator)
		srv.GET("/version", serviceVersionGet)
	}

	r.Run(":8080")
}
