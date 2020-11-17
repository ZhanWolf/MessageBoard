package controller

import (
	dao2 "MessageBoard/dao"
	"MessageBoard/service"
	"MessageBoard/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {

}

func (uc *UserController) Router(engine *gin.Engine)  {
	engine.POST("/registe", uc.registe)
}

//用户注册
func (uc *UserController) registe(ctx *gin.Context) {
	userName := ctx.PostForm("username")
	pwd := ctx.PostForm("password")
	fmt.Println("user:",  userName,  pwd)

	thisDao := dao2.UserDao{tool.GetDb()}
	flag := thisDao.QueryUsername(userName)
	if flag == true {
		ctx.JSON(200, gin.H{
			"message": "该用户已经存在！",
		})
		return
	}

	us := new(service.UserService)
	ok := us.RegisteByPwd(userName, pwd)
	if ok == true {
		ctx.JSON(200, gin.H{
			"message": "注册成功！",
			"username": userName,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "注册失败！",
	})

}
