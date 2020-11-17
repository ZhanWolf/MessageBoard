package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type HelloController struct {

}

func (hc *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hc.Hello)
}

//输出HelloWorld进行调试
func (hc *HelloController) Hello(ctx *gin.Context)  {
	ctx.JSON(200, gin.H{
		"message": "HelloWorld!",
	})

	fmt.Println("HelloWorld!")
}