package tool

import (
	"github.com/gin-gonic/gin"
)

//检查登录状态，无cookie返回""，否则返回用户名
func CheckLogin(ctx *gin.Context) string {
	cookie, err := ctx.Request.Cookie("isLogin")
	if err != nil {
		return ""
	}

	value := cookie.Value
	return value
}
