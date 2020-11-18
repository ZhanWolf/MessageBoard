package controller

import (
	"MessageBoard/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {

}

func (uc *UserController) Router(engine *gin.Engine)  {
	engine.POST("/registe", uc.registe)
	engine.POST("/login", uc.login)
}

//用户登录
func (uc *UserController) login(ctx *gin.Context) {
	userName := ctx.PostForm("username")
	pwd := ctx.PostForm("password")
	fmt.Println("LoginUserInfo:", userName, pwd)

	us := new(service.UserService)

	flag := us.CheckUserAlive(userName)
	if flag == true {
		ctx.JSON(200, gin.H{
			"message": "该用户不存在！",
		})
		return
	}

	cookie := us.LoginByPwd(userName, pwd)
	if cookie == nil {
		ctx.JSON(200, gin.H{
			"message": "密码错误！",
		})
		return
	}

	http.SetCookie(ctx.Writer, cookie)
	ctx.JSON(200, gin.H{
		"message": "登录成功！",
	})
}

//用户注册
func (uc *UserController) registe(ctx *gin.Context) {
	userName := ctx.PostForm("username")
	pwd := ctx.PostForm("password")
	fmt.Println("RegisteUserInfo:", userName, pwd)

	us := new(service.UserService)

	flag := us.CheckUserAlive(userName)
	if flag == false {
		ctx.JSON(200, gin.H{
			"message": "该用户已经存在！",
		})
		return
	}

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
