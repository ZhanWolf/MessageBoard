package dao

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"time"
)

type MessageDao struct {
	*sql.DB
}

func (md *MessageDao) ListMsgs(ctx *gin.Context) error {
	rows, err := md.Query("select * from message_info")
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var message string
		var username string
		var time int64
		err = rows.Scan(&id, &message, &username, &time)
		if err != nil {
			return err
		}

		ctx.JSON(200, gin.H{
			"id": id,
			"message": message,
			"username": username,
			"time": time,
		})
	}
	err = rows.Err()
	return err
}

//删除数据库中的一条留言
func (md *MessageDao) DeleteMsg(id string) error {
	_, err := md.Exec("delete from message_info where id = ?", id)
	return err
}

//向数据库中插入一条信息
func (md *MessageDao) InsertMessage(message, username string) error {
	insertTime := time.Now().Unix()
	_, err := md.Exec(" insert into message_info(message, username, time) " + "value(?, ?, ?)", message, username, insertTime)
	return err
}
