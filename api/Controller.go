package api

import (
	m "WorkaholicSrv/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}

func AddRecordApi(c *gin.Context) {
	username := c.Request.FormValue("name")
	timestamp, err := strconv.ParseInt(c.Request.FormValue("start_time"), 10, 64)
	recordtype, err := strconv.Atoi(c.Request.FormValue("type"))
	desc := c.Request.FormValue("description")

	or := m.OvertimeRecord{UserName: username, StartTimestamp: timestamp, RecordType: recordtype, Description: desc}
	err = or.AddOvertimeRecord()
	checkErr(err)

	c.JSON(http.StatusOK, gin.H{
		"msg": "Insert Success",
	})
}

func ModifyRecordApi(c *gin.Context) {
	username := c.Request.FormValue("name")
	timestamp, err := strconv.ParseInt(c.Request.FormValue("finish_time"), 10, 64)
	recordtype, err := strconv.Atoi(c.Request.FormValue("type"))

	or := m.OvertimeRecord{UserName: username, FinishTimestamp: timestamp, RecordType: recordtype}
	err = or.AddOvertimeRecord()
	checkErr(err)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Insert Success",
	})
}

func GetRecordsApi(c *gin.Context) {
	username := c.Request.FormValue("name")
	or := m.OvertimeRecord{UserName: username}
	records, err := or.GetOvertimeRecords()
	checkErr(err)
	c.JSON(http.StatusOK, gin.H{
		"records": records,
	})
}
