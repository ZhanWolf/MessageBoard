package dao

import (
	"database/sql"
	"time"
)

type UserDao struct {
	*sql.DB
}

//根据用户名查询数据库中用户是否已经存在
func (dao *UserDao) QueryUsername(username string) bool{
	row, err := dao.Query("select username from userinfo")
	var str string
	err = row.Scan(str)
	if err != nil {
		panic(err.Error())
	}

	if str != "" {
		return true
	}
	return false
}

func (dao *UserDao) InsertUser(username, pwd string) error {
	registerTime := time.Now().Unix()
	_, err := dao.Exec("insert into userinfo(username, password, register_time) " + "values(?, ?, ?);", username, pwd, registerTime)
	return err
}