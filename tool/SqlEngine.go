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
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/message_board")
	if err != nil {
		fmt.Println(err)
		return
	}
	Db = db
//	fmt.Println("FSfdsasdf", db)
}
