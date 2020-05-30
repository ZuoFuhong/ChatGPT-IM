package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-IM/config"
	"log"
)

var Cli *sql.DB

func init() {
	conf := config.LoadConf().Mysql
	addr := fmt.Sprintf("%s:%s@tcp(%s)/gim?charset=utf8&parseTime=true", conf.Username, conf.Password, conf.Addr)

	var err error
	Cli, err = sql.Open("mysql", addr)
	if err != nil {
		log.Panic(err)
	}
}
