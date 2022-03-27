package main

import (
	"gin_demo2/common"
	"gin_demo2/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := common.GetDB()
	defer db.Close()
	r := gin.Default()
	r = routes.CollectRoute(r)

	panic(r.Run())
}
