package main

import (
	"log"
	"main/DB"
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
}
