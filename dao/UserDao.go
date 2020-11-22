package dao

import (
	"database/sql"
	"fmt"
	"time"
)

type UserDao struct {
	*sql.DB
}

func (dao *UserDao) ChangePwd(usernama, newPwd string) error {
	_, err := dao.Exec("update userinfo set password = ? where username = ?", newPwd, usernama)
	return err
}

//根据用户名返回用户密码
func (dao *UserDao) QueryUserPwd(username string) string {
	row := dao.QueryRow("select password from userinfo where username = ? ", username)

	err := row.Err()
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	var pwd string
	row.Scan(&pwd)
	return pwd
}

//根据用户名查询数据库中用户是否已经存在
func (dao *UserDao) QueryUsername(name string) bool{

	row := dao.QueryRow("select username from userinfo where username = ? ", name)

	err := row.Err()
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	var usr string
	row.Scan(&usr)
	if usr == "" {
		return true
	}

	fmt.Println(usr)
	return false
}

//插入注册信息
func (dao *UserDao) InsertUser(username, pwd string) error {
	registerTime := time.Now().Unix()
	_, err := dao.Exec("insert into userinfo(username, password, register_time) " + "values(?, ?, ?);", username, pwd, registerTime)
	return err
}