package api

import (
	m "WorkaholicSrv/model"
	"fmt"
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
	timestamp, err := strconv.ParseInt(c.Request.FormValue("time"), 10, 64)
	recordtype, err := strconv.Atoi(c.Request.FormValue("type"))

	fmt.Println(username)
	fmt.Println(timestamp)
	fmt.Println(recordtype)
	or := m.OvertimeRecord{UserName: username, Timestamp: timestamp, RecordType: recordtype}
	ra, err := or.AddRecord()
	checkErr(err)
	msg := fmt.Sprintf("Insert success %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
