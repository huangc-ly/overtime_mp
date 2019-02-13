package database

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
		return
	}
}

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/workaholic?charset=utf8")
	CheckErr(err)
	SqlDB.SetMaxIdleConns(20)
	SqlDB.SetMaxOpenConns(20)
	err = SqlDB.Ping()
	CheckErr(err)
}
