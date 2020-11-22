package main

import (
	"MessageBoard/tool"
	"github.com/gin-gonic/gin"
	"MessageBoard/controller"
)

func main() {
	tool.SqlEngine()

	engin := gin.Default()
	
	registRouter(engin)
	
	engin.Run()
}

//调用controller
func registRouter(engine *gin.Engine)  {
	new(controller.HelloController).Router(engine)
	new(controller.UserController).Router(engine)
	new(controller.MessageController).Router(engine)
}