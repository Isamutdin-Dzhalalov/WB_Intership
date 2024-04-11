package main

import (
	"log"
//	"os"
	"main/DB"
//	"main/setting"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
//	stan "github.com/nats-io/stan.go"

//	"io/ioutil"
// 	"fmt"
//	"encoding/json"
)

var router *gin.Engine

func main() {

	db, err := DB.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}


	tables := []string{"order_meta", "delivery", "payment", "item"}
	DB.DropTable(db, tables)
	DB.CreateTables(db)
	DB.InsertDataInTable()

/*
	sc, err := stan.Connect("test-cluster", "test-client", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal("stan.Connect: ", err)
	}
	defer sc.Close()

	_, err = sc.Subscribe("foo", func(msg *stan.Msg) {
		DB.InsertDataInTable()
	}, stan.StartWithLastReceived())
//	select{}
*/

//	tables := []string{"order_meta", "delivery", "payment", "item"}
//	fmt.Printf("Order: %+v\n", order)
//	DB.DropTable(db, tables)
//	DB.CreateTables(db)
//	DB.PrintTable(db, "order_meta")
//	DB.InsertDataInTable()

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
