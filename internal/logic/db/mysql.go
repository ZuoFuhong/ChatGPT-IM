package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Cli *sql.DB

func init() {
	var err error
	Cli, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/gim?charset=utf8&parseTime=true")
	if err != nil {
		log.Panic(err)
	}
}
