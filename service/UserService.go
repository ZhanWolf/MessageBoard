package service

import (
	dao2 "MessageBoard/dao"
	"MessageBoard/tool"
	"fmt"
	"net/http"
)

type UserService struct {

}

//登录服务
func (us *UserService) LoginByPwd(username, pwd string) *http.Cookie  {
	thisDao := dao2.UserDao{tool.GetDb()}
	rightPwd := thisDao.QueryUserPwd(username)

	if rightPwd != pwd {
		return nil
	}

	cookie := &http.Cookie{
		Name: username,
		Value: pwd,
		Path: "/",
		HttpOnly: true,
	}
	return cookie
}

//检查用户是否存在, true不存在, false存在
func (us *UserService) CheckUserAlive(username string) bool {
	thisDao := dao2.UserDao{ tool.GetDb() }
	return thisDao.QueryUsername(username)
}

//注册服务
func (us *UserService) RegisteByPwd(userName, pwd string) bool {
	//插入数据库
	thisDao := dao2.UserDao{tool.GetDb()}
	err := thisDao.InsertUser(userName, pwd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
