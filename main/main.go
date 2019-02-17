package main

import (
	. "WorkaholicSrv/api"
	db "WorkaholicSrv/database"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/AddRecord", AddRecordApi)
	router.POST("/ModifyRecord", ModifyRecordApi)
	router.GET("/GetRecords", GetRecordsApi)
	return router
}

func main() {
	defer db.SqlDB.Close()
	r := initRouter()
	r.Run(":12000") // listen and serve on 0.0.0.0:8080
}
