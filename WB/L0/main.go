package main

import (
	"log"
	"main/DB"
)

func main() {
	db, err := DB.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

//	tables := []string{"order_meta", "delivery", "payment", "item"}
//	DB.DropTable(db, tables)

	DB.CreateTables(db)
	DB.InsertDataInTable(db)
}
