package tool

import "github.com/gin-gonic/gin"

func PrintInfo(ctx *gin.Context, info interface{})  {
	ctx.JSON(200, gin.H{
		"data": info,
	})
}
