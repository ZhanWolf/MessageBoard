package controller

import (
	"MessageBoard/service"
	"MessageBoard/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type MessageController struct {

}

func (mc *MessageController) Router(engine *gin.Engine) {

	engine.POST("/msgs", mc.SendMsg)
	engine.DELETE("/msgs/:id", AdminMiddleWare(), mc.deleteMsg)
	engine.GET("/msgs", mc.listMsgs)
}

//列出所有留言
func (mc *MessageController) listMsgs(ctx *gin.Context) {
	ms := service.MessageService{}
	err := ms.ListMsgs(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//删除一条留言
func (mc *MessageController) deleteMsg(ctx *gin.Context) {
	id := ctx.Param("id")
	ms := service.MessageService{}
	err := ms.DeleteMsg(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	tool.PrintInfo(ctx, "删除成功！")
}

//新建一条留言
func (mc *MessageController) SendMsg(ctx *gin.Context) {
	username := tool.CheckLogin(ctx)
	if username == "" {
		tool.PrintInfo(ctx, "先登录在进行操作 ")
		return
	}

	message := ctx.PostForm("message")
	ms := service.MessageService{}
	err := ms.SendMsg(message, username)
	if err != nil {
		fmt.Println(err)
		return
	}

	 tool.PrintInfo(ctx, "发表留言成功 ")
}

//管理员权限中间件
func AdminMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("isLogin")
		if err == nil {
			username := cookie.Value
			if username == "wmf" {
				ctx.Next()
				return
			}
		}
		tool.PrintInfo(ctx, "你不是管理员！")
		ctx.Abort()
		return
	}
}