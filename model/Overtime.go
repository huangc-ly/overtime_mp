package model

import (
	db "WorkaholicSrv/database"
	"log"
	"fmt"
)

type OvertimeRecord struct {
	UserName   string `json:"name"`
	Timestamp  int64  `json:"time"`
	RecordType int    `json:"type"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
		return
	}
}

func (p *OvertimeRecord) AddRecord() (id int64, err error) {

	if db.SqlDB == nil {
		fmt.Println("db.SqlDB nil")
		return
	}
	stmt, err := db.SqlDB.Prepare("INSERT INTO work(user_name, time) VALUES (?, ?)")
	if p.RecordType == 1 {
		rs, err := stmt.Exec(p.UserName, p.Timestamp)
		checkErr(err)
		id, err = rs.LastInsertId()
	} else if p.RecordType == 0 {
		rs, err := db.SqlDB.Exec("INSERT INTO offwork(user_name, time) VALUES (?, ?)", p.UserName, p.Timestamp)
		checkErr(err)
		id, err = rs.LastInsertId()
	}
	return
}
