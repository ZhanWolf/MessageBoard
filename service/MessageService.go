package service

import (
	dao2 "MessageBoard/dao"
	"MessageBoard/tool"
	"github.com/gin-gonic/gin"
)

type MessageService struct {

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
