package main

import (
	. "WorkaholicSrv/api"
	db "WorkaholicSrv/database"

	"fmt"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/record", AddRecordApi)
	return router
}

func main() {
	if db.SqlDB == nil {
		fmt.Println("db.SqlDB nil")
	}
	defer db.SqlDB.Close()
	r := initRouter()
	r.Run(":12000") // listen and serve on 0.0.0.0:8080
}
