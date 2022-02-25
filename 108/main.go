package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Mac       string `json:"mac" binding:"mac"`
	Email     string `json:"email" binding:"email"`
	Ascii     string `json:"ascii" binding:"min=6,max=10,printascii"`
	Number    int    `json:"number" binding:"required"`
	NumberPtr *int   `json:"numberPtr" binding:"required"`
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

func serviceGetValidator(c *gin.Context) {
	var valid Validator
	var a int
	a = 12
	valid.NumberPtr = &a
	fmt.Printf("valid: %+v\n", valid)

	c.JSON(http.StatusOK, valid)
}

func serviceVersionGet(c *gin.Context) {
	c.Status(http.StatusOK)
}

func Validate(s interface{}) error {
	validate := validator.New()
	if errs := validate.Struct(s); errs != nil {
		fmt.Println(errs)
		return errs
	} else {
		fmt.Println(errs)
	}
	return nil
}
func main() {

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 服务类
	srv := r.Group("/service")
	{
		srv.PUT("/validator", servicePutValidator)
		srv.GET("/validator", serviceGetValidator)
		srv.GET("/version", serviceVersionGet)
	}

	r.Run(":10100")
}
