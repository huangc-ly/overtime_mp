package model

import (
	db "WorkaholicSrv/database"
	"log"
)

type OvertimeRecord struct {
	UserName        string `json:"name"`
	StartTimestamp  int64  `json:"start_time"`
	RecordType      int    `json:"type"`
	FinishTimestamp int64  `json:"finish_time"`
	Description     string `json:"description"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
		return
	}
}

func (p *OvertimeRecord) AddOvertimeRecord() (err error) {

	if p.RecordType == 1 { //插入一条加班数据
		_, err := db.SqlDB.Exec("INSERT INTO overtime_record(user_name, start_time, description, status, finish_time) VALUES (?, ?, ?, ?, ?)", p.UserName, p.StartTimestamp, p.Description, 1, 0)
		checkErr(err)
	} else if p.RecordType == 0 { //首先找到用户当前未关闭加班记录
		rows, err := db.SqlDB.Query("SELECT id FROM overtime_record WHERE user_name=? AND status=?", p.UserName, 1)
		checkErr(err)
		defer rows.Close()
		var id int64
		for rows.Next() {
			err = rows.Scan(&id)
			checkErr(err)
			_, err := db.SqlDB.Exec("UPDATE overtime_record SET finish_time=?, status=? WHERE id=?", p.FinishTimestamp, 0, id)
			checkErr(err)
		}
	}

	return
}

func (p *OvertimeRecord) GetOvertimeRecords() (records []OvertimeRecord, err error) {

	records = make([]OvertimeRecord, 0)

	rows, err := db.SqlDB.Query("SELECT id FROM overtime_record WHERE user_name=? AND status=?", p.UserName, 0)
	checkErr(err)
	defer rows.Close()
	var r OvertimeRecord
	for rows.Next() {
		err = rows.Scan(&r.StartTimestamp, &r.FinishTimestamp, &r.Description)
		checkErr(err)
		records = append(records, r)
	}
	return
}
