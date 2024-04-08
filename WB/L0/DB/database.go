package DB

import (
	"fmt"
	"database/sql"
	"log"
	"os"
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
		name VARCHAR(30) NOT NULL 
	);`

	_, err := db.Exec(createOrder)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}


/*
	createOrder := `
	CREATE TABLE IF NOT EXISTS order_meta (
		order_uid VARCHAR(24) PRIMARY KEY,
		track_number VARCHAR(255) NOT NULL,
		entry VARCHAR(255) NOT NULL,
		locale VARCHAR(255),
		internal_signature VARCHAR(255),
		customer_id VARCHAR(255) NOT NULL,
		delivery_service VARCHAR(255),
		shardkey VARCHAR(255),
		sm_id INT,
		date_created VARCHAR(255),
		oof_shard VARCHAR(255)
	);`

	createDelivery := `
	CREATE TABLE IF NOT EXISTS delivery (
		order_uid VARCHAR(24) PRIMARY KEY,
		FOREIGN KEY (order_uid) REFERENCES order_meta(order_uid),
		data_delivery JSON NOT NULL
	);`

		createPayment := `
	CREATE TABLE IF NOT EXISTS payment (
		order_uid VARCHAR(24) PRIMARY KEY,
		FOREIGN KEY (order_uid) REFERENCES order_meta(order_uid),
		data_payment JSON NOT NULL
	);`

	createItem := `
	CREATE TABLE IF NOT EXISTS item (
		id_item INTEGER PRIMARY KEY,
		data_item JSON NOT NULL
	);`

	createOrderItem := `
	CREATE TABLE IF NOT EXISTS order_item (
		id_cart VARCHAR(24) REFERENCES order_meta(order_uid),
		id_item INTEGER REFERENCES item(id_item),
		PRIMARY KEY (id_cart, id_item)
	);`

//	tables := []string{"createOrder", "createDelivery", "createItem", "createPayment", "createOrderItem"}
//	for _, table := range tables {
//		_, err := db.Exec(table)
//		if err != nil {
//			log.Fatal("CReateTables func :", err)
//		}
//	}
	_, err := db.Exec(createOrder)

	if err != nil {
		log.Fatal("CreateTabcles func :", err)
	}

	_, err = db.Exec(createDelivery)

	if err != nil {
		log.Fatal("CreateTabcles func :", err)
	}
	
	_, err = db.Exec(createItem)

	if err != nil {
		log.Fatal("CreateTabcles func :", err)
	}
	
	_, err = db.Exec(createPayment)

	if err != nil {
		log.Fatal("CreateTabcles func :", err)
	}

	_, err = db.Exec(createOrderItem)

	if err != nil {
		log.Fatal("CreateTabcles func :", err)
	}
*/
}

func DropTable(db *sql.DB, tableName string) {
//	dropTable := "DROP TABLE IF EXISTS %s;"
	dropTable := "DROP TABLE IF EXISTS %s CASCADE;"
	query := fmt.Sprintf(dropTable, tableName)

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("DropTable: ", err)
	}
}

func PrintTable(db *sql.DB, tableName string) {
	sqlRequest := "SELECT * FROM %s "
	query := fmt.Sprintf(sqlRequest, tableName)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("PrintTable func -> rows ", err)
		os.Exit(1)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("PrintTable func -> columns: ", err)
	}

		// Создание слайса для хранения значений каждой строки
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))

	// Заполнение слайса значениями для сканирования
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Вывод данных
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal("PrintTable func -> for cicle :", err)
		}

		// Вывод значений каждой строки
		var value string
		for i, col := range values {
			if col != nil {
				value = fmt.Sprintf("%s", col)
			} else {
				value = "NULL"
			}
			fmt.Printf("%s: %s\n", columns[i], value)
		}
		fmt.Println("---")
	}

	if err = rows.Err(); err != nil {
		log.Fatal("PrintTable func -> rows.Err: ", err)
	}
}


