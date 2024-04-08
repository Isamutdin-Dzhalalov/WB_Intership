package main

import (
	"log"
	"os"
	"main/DB"
	"main/setting"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var router *gin.Engine

func main() {
	_, err := DB.ConnectWithDb()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

/*
	DB.DropTable(db, "order_meta")
	DB.CreateTables(db)
	_, err = db.Exec(`INSERT INTO "order_meta" ("name") VALUES('AAAAAAQQQ')`)
	if err != nil {
	log.Fatal("main -> db.Exec: ", err)
		os.Exit(1)
	}
	DB.PrintTable(db, "order_meta")
*/

	router = gin.Default()
	router.Static("/qwe", setting.Config.HTML)
	router.LoadHTMLFiles(setting.Config.HTML + "index.html")
	router.GET("/", index)
	router.Run(setting.Config.ServerHost + ":" + setting.Config.ServerPort)
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
