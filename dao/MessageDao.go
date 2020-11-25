package dao

import (
	"MessageBoard/Struct"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type MessageDao struct {
	*sql.DB
}

//递归处理节点
func (md *MessageDao) SolveSon(ctx *gin.Context, myInfo *Struct.Info) error {
	//进入时自己已输出信息, 自己的信息存在myInfo中
	rows, err := md.Query("select id, message, username, time, comment_num from message_info where pid = ?", myInfo.Id)
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		//获得每个儿子的信息
		var message string
		var username string
		var time int64
		var commentNum int
		var id string
		err := rows.Scan(&id, &message, &username, &time, &commentNum)
		if err != nil {
			return err
		}

		//把儿子的信息存在结构体中
		sonInfo := new(Struct.Info)
		sonInfo.Id, _ = strconv.Atoi(id)
		sonInfo.Username = username
		sonInfo.Message = message

		//儿子和爸爸的id一样，排除该情况
		if sonInfo.Id == myInfo.Id {
			continue
		}

		//输出信息
		ctx.JSON(200, gin.H{
			"originMsg": myInfo.Message,
			"originUser": myInfo.Username,
			"originId": myInfo.Id,
			"commentMsg": message,
			"commentUser": username,
			"commentId": id,
			"commentTime": time,
			"commentNum": commentNum,
		})

		//递归处理
		err = md.SolveSon(ctx, sonInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

//返回一个节点信息
func (md *MessageDao) GetInfo(id string) (*Struct.Info, error) {
	row := md.QueryRow("select message, username from message_info where id = ?", id)

	var message string
	var username string
	err := row.Scan(&message, &username)
	if err != nil {
		return nil, err
	}

	Info := new(Struct.Info)
	Info.Id, _ = strconv.Atoi(id)
	Info.Message = message
	Info.Username = username

	return Info, nil
}

//获取一条信息
func (md *MessageDao) GetOneMsgs(ctx *gin.Context, id string) error {
	row := md.QueryRow("select message, username, time, comment_num from message_info where id = ?", id)

	var message string
	var username string
	var time int64
	var commentNum int
	err := row.Scan(&message, &username, &time, &commentNum)
	if err != nil {
		return err
	}

	ctx.JSON(200, gin.H{
		"id": id,
		"message": message,
		"username": username,
		"time": time,
		"comment_num": commentNum,
	})

	err = row.Err()
	return err
}


func (md *MessageDao) ListMsgs(ctx *gin.Context) error {
	rows, err := md.Query("select id, message, username, time, comment_num from message_info where id = pid")
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var message string
		var username string
		var time int64
		var commentNum int
		err = rows.Scan(&id, &message, &username, &time, &commentNum)
		if err != nil {
			return err
		}

		ctx.JSON(200, gin.H{
			"id": id,
			"message": message,
			"username": username,
			"time": time,
			"comment_num": commentNum,
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
	if err != nil {
		return err
	}

// 设置parentId为id，但是这么写会报错，You can't specify target table 'message_info' for update in FROM clause，只能搞麻烦一点
//	_, err = md.Exec("update message_info set pid = (select id from message_info where username = ?) where username = ?", username, username)
//	return err

	//先获取id
	var id int
	row := md.QueryRow("SELECT LAST_INSERT_ID()")
	row.Scan(&id)
	err = row.Err()
	if err != nil {
		return err
	}

	//再设置父节点为id
	_, err = md.Exec("update message_info set pid = ? where id = ?", id, id)
	return err
}

//向数据库中插入一条评论
func (md *MessageDao) InsertComment(pid, username, message string) error {
	insertTime := time.Now().Unix()
	_, err := md.Exec(" insert into message_info(message, username, time, pid) " + "value(?, ?, ?, ?)", message, username, insertTime, pid)
	if err != nil {
		return err
	}

	//把爸爸的评论数+1s
	var commentNum int
	row := md.QueryRow("select comment_num from message_info where id = ?", pid)
	err = row.Err()
	if err != nil {
		return err
	}

	err = row.Scan(&commentNum)
	if err != nil {
		return err
	}

	commentNum++
	_, err = md.Exec("update message_info set comment_num = ? where id = ?", commentNum, pid)
	if err != nil {
		return err
	}

	return nil
}
