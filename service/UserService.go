package service

import (
	dao2 "MessageBoard/dao"
	"MessageBoard/tool"
	"log"
)

type UserService struct {

}

//注册操作
func (us *UserService) RegisteByPwd(userName, pwd string) bool {
	//插入数据库
	thisDao := dao2.UserDao{tool.GetDb()}
	err := thisDao.InsertUser(userName, pwd)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
