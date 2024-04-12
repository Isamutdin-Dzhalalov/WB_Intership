package main

import (
//	"fmt"
	"log"
	"main/DB"
//	"main/cache"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	  "github.com/rainycape/memcache"
)

var router *gin.Engine
var mc *memcache.Client

func main() {

	db, err := DB.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}


	tables := []string{"order_meta", "delivery", "payment", "item"}
	DB.DropTable(db, tables)
	DB.CreateTables(db)
	DB.InsertDataInTable()
//	DB.PrintTable(db, "order_meta")
//	cache.CacheData()
}	
/*

	 Работа с html

	router = gin.Default()
	router.Static("/qwe", setting.Config.HTML)
	router.LoadHTMLFiles(setting.Config.HTML + "index.html")
	router.GET("/", index)
	router.Run(setting.Config.ServerHost + ":" + setting.Config.ServerPort)
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
*/
