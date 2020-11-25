package service

import (
	"MessageBoard/Struct"
	dao2 "MessageBoard/dao"
	"MessageBoard/tool"
	"github.com/gin-gonic/gin"
)

type MessageService struct {

}

func (ms *MessageService) GetInfo(id string) (*Struct.Info, error) {
	thisDao := dao2.MessageDao{tool.GetDb()}
	Info, err := thisDao.GetInfo(id)
	return Info, err
}

//输出套娃评论
func (ms *MessageService) TaoWa(ctx *gin.Context, info *Struct.Info) error {
	thisDao := dao2.MessageDao{tool.GetDb()}
	err := thisDao.SolveSon(ctx, info)
	return err
}

//列出一条留言
func (ms *MessageService) GetOneMsgs(ctx *gin.Context, id string) error {
	thisDao := dao2.MessageDao{tool.GetDb()}
	err := thisDao.GetOneMsgs(ctx, id)
	return err
}

//列出所有留言
func (ms *MessageService) ListMsgs(ctx *gin.Context) error {
	thisDao := dao2.MessageDao{tool.GetDb()}
	err := thisDao.ListMsgs(ctx)
	return err
}

//删除留言
func (ms *MessageService) DeleteMsg(id string) error {
	thisDao := dao2.MessageDao{tool.GetDb()}
	err := thisDao.DeleteMsg(id)
	return err
}

//发送留言
func (ms *MessageService) SendMsg(message, username string) error {
	thisDao := dao2.MessageDao{ tool.GetDb() }
	err := thisDao.InsertMessage(message, username)
	return err
}

//对留言或者评论，发送评论
func (ms *MessageService) SendComment(pid, username, message string) error {
	thisDao := dao2.MessageDao{tool.GetDb()}
	err := thisDao.InsertComment(pid, username, message)
	return err
}
