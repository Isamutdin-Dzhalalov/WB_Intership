package DB

import (
	"fmt"
	"database/sql"
//	"encoding/json"
	"log"
	"os"
//	"io/ioutil"
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
		  amount INT,
		  payment_dt BIGINT,
		  bank VARCHAR(50),
		  delivery_cost INT,
		  goods_total INT,
		  custom_fee INT
	);`
	

		createItem := `
	CREATE TABLE IF NOT EXISTS item (
		 chrt_id BIGSERIAL PRIMARY KEY,
		 order_uid VARCHAR(50) REFERENCES order_meta(order_uid),
		 track_number VARCHAR(30),
		 price INT,
		 rid VARCHAR(50),
		 name VARCHAR(30),
		 sale INT,
		 size VARCHAR(30),
		 total_price INT,
		 nm_id INT,
		 brand VARCHAR(30),
		 status INT
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
		log.Fatal("CreateTabcles createOrder :", err)
	}

	_, err = db.Exec(createDelivery)
	if err != nil {
		log.Fatal("CreateTabcles createDelivery :", err)
	}
	
	_, err = db.Exec(createItem)

	if err != nil {
		log.Fatal("CreateTabcles createItem :", err)
	}
	
	_, err = db.Exec(createPayment)

	if err != nil {
		log.Fatal("CreateTabcles createPayment :", err)
	}
}

func DropTable(db *sql.DB, tables []string) {
//	dropTable := "DROP TABLE IF EXISTS %s;"
	dropTable := "DROP TABLE IF EXISTS %s CASCADE;"
	for _, table := range tables {
		query := fmt.Sprintf(dropTable, table)
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
	
/*
	query := fmt.Sprintf(dropTable, tableName)

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("DropTable: ", err)
	}
*/
}


/*
func DecodeJsonToStruct() {
	fileData, err := ioutil.ReadFile("model.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileData)
	var root Root
	err = json.Unmarshal([]byte(fileData), &root)
	if err != nil {
		log.Fatal("Unmarshal :", err)
	}
	fmt.Println(root)
	file, err := os.Open("model.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileJsonValue, _ := ioutil.ReadAll(file)
	fmt.Println(string(fileJsonValue))
}
*/

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


