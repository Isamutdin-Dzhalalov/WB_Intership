package main

import (
	"log"
	"os"
	"main/DB"
//	"main/setting"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

//	"io/ioutil"
//	"fmt"
//	"encoding/json"
)

var router *gin.Engine

func main() {

	db, err := DB.ConnectDB()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	tables := []string{"order_meta", "delivery", "payment", "item"}
//	fmt.Printf("Order: %+v\n", order)

	DB.DropTable(db, tables)
	DB.CreateTables(db)
//	DB.PrintTable(db, "order_meta")
	DB.InsertDataInTable()

//	return 

//	DB.DecodeJsonToStruct()

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
