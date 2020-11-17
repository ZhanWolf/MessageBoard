package tool

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func GetDb() *sql.DB {
	return Db
}

func SqlEngine()  {
	db, err := sql.Open("message_board", "root:12345678@tcp(localhost:3306)/ginsql")
	if err != nil {
		fmt.Println(err)
	}
	Db = db
}
