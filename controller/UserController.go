package controller

import (
	"MessageBoard/service"
	"MessageBoard/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {

}

func (uc *UserController) Router(engine *gin.Engine)  {
	engine.POST("/registe", uc.registe)
	engine.POST("/login", uc.login)
	engine.POST("/changePwd", uc.changePwd)
	engine.GET("/logout", uc.logout)
}


//退出登录
func (uc *UserController) logout(ctx *gin.Context) {
	value := tool.CheckLogin(ctx)
	if value == "" {
		tool.PrintInfo(ctx, "未登录 ")
		return
	}

	cookie, err := ctx.Request.Cookie("isLogin")
	if err != nil {
		tool.PrintInfo(ctx, "获取cookie失败")
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(ctx.Writer, cookie)

	tool.PrintInfo(ctx, "退出登录成功")
}

//修改密码
func (uc *UserController) changePwd(ctx *gin.Context) {
	//验证登录状态，只有登录才能修改密码
	username := tool.CheckLogin(ctx)
	newPwd := ctx.PostForm("newPwd")
	if username == "" {
		tool.PrintInfo(ctx, "请先登录 ")
		return
	}
	//service
	us := service.UserService{}
	err := us.ChangePwd(username, newPwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(username, "newPwd: ", newPwd)
	tool.PrintInfo(ctx, "修改密码成功")
}

//用户登录
func (uc *UserController) login(ctx *gin.Context) {
	value := tool.CheckLogin(ctx)
	if value != "" {
		tool.PrintInfo(ctx, "用户已登录 ")
		return
	}

	userName := ctx.PostForm("username")
	pwd := ctx.PostForm("password")
	fmt.Println("LoginUserInfo:", userName, pwd)

	us := new(service.UserService)

	flag := us.CheckUserAlive(userName)
	if flag == true {
		tool.PrintInfo(ctx, "该用户不存在 ")
		return
	}

	cookie := us.LoginByPwd(userName, pwd)
	if cookie == nil {
		tool.PrintInfo(ctx, "密码错误 ")
		return
	}

	http.SetCookie(ctx.Writer, cookie)
	tool.PrintInfo(ctx, "登录成功 ")
}

//用户注册
func (uc *UserController) registe(ctx *gin.Context) {
	userName := ctx.PostForm("username")
	pwd := ctx.PostForm("password")
	fmt.Println("RegisteUserInfo:", userName, pwd)

	us := new(service.UserService)

	flag := us.CheckUserAlive(userName)
	if flag == false {
		tool.PrintInfo(ctx, "该用户已经存在 ")
		return
	}

	ok := us.RegisteByPwd(userName, pwd)
	if ok == true {
		tool.PrintInfo(ctx, "注册成功 ")
		return
	}
	tool.PrintInfo(ctx, "注册失败 ")

}
