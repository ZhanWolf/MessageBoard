package model

type Message struct {
	Id int `form:"id"`
	Username int `form:"username"`
	Time int64 `form:"time"`
	Message string `form:"message"`
}