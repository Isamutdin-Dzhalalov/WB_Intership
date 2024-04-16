package DB

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

type Config struct {
	User string
	Password string
	Dbname string
}

func CreateTables(db *sql.DB) {

	createOrder := `
	CREATE TABLE IF NOT EXISTS order_meta (
		order_uid VARCHAR(240) PRIMARY KEY,
		track_number VARCHAR(255) NOT NULL,
		entry VARCHAR(255) NOT NULL,
		locale VARCHAR(255),
		internal_signature VARCHAR(255),
		customer_id VARCHAR(255) NOT NULL,
		delivery_service VARCHAR(255),
		shardkey VARCHAR(255),
		sm_id BIGINT,
		date_created VARCHAR(255),
		oof_shard VARCHAR(255)
	);`

	createDelivery := `
	CREATE TABLE IF NOT EXISTS delivery (
	   order_uid VARCHAR(24) REFERENCES order_meta(order_uid),
	   name VARCHAR(30),
	   phone VARCHAR(20),
	   zip VARCHAR(100),
	   city VARCHAR(30),
	   address VARCHAR(100),
	   region VARCHAR(50),
	   email VARCHAR(30)
	);`

		createPayment := `
	CREATE TABLE IF NOT EXISTS payment (
		order_uid VARCHAR(24) REFERENCES order_meta(order_uid),
		  transaction VARCHAR(50),
		  request_id VARCHAR(30),
		  currency VARCHAR(50),
		  provider VARCHAR(50),
		  amount BIGINT,
		  payment_dt BIGINT,
		  bank VARCHAR(50),
		  delivery_cost BIGINT,
		  goods_total BIGINT,
		  custom_fee BIGINT
	);`
	

		createItem := `
	CREATE TABLE IF NOT EXISTS item (
		 chrt_id BIGSERIAL PRIMARY KEY,
		 order_uid VARCHAR(50) REFERENCES order_meta(order_uid),
		 track_number VARCHAR(30),
		 price BIGINT,
		 rid VARCHAR(50),
		 name VARCHAR(30),
		 sale BIGINT,
		 size VARCHAR(30),
		 total_price BIGINT,
		 nm_id BIGINT,
		 brand VARCHAR(30),
		 status BIGINT
	);`

	tables := []string{createOrder, createDelivery, createItem, createPayment}
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			log.Fatal("CReateTables func -> table:", err)
		}
	}
}

func DropTable(db *sql.DB, tables []string) {

	dropTable := "DROP TABLE IF EXISTS %s CASCADE;"
	for _, table := range tables {
		query := fmt.Sprintf(dropTable, table)
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
}
