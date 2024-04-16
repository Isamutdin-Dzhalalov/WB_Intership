package DB

import (
	"log"
	"encoding/json"
	stan "github.com/nats-io/stan.go"
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
)

const (
	insertOrderMeta = `
		INSERT INTO order_meta (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, 
		shardkey, sm_id, date_created, oof_shard)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

	insertDelivery = `
		INSERT INTO delivery (name, phone, zip, city, address, region, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7);`

	insertPayment = `
		INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`

	insertItem = `
		INSERT INTO item (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`
)

func Error(err error, str string) {
	if err != nil {
		log.Fatal("InsertDataInTable -> ", str, err)
	}
}

func InsertDataInTable(db *sql.DB) {

	sc, err := stan.Connect("test-cluster", "test-client")
	Error(err, "stan.Connect: ")

	defer sc.Close()

	file := DataGeneration()

	err = sc.Publish("foo", file)

	var cache *Cache
	_, err = sc.Subscribe("foo", func(msg *stan.Msg) {
		var order Order
		err = json.Unmarshal([]byte(file), &order)
		Error(err, "json.Unmarshal: ")
	
		db, err := ConnectDB()
		Error(err, "ConnectDB: ")

		orDel := order.Delivery
		orPay := order.Payment
		orItem := order.Items[0]

		_, err = db.Exec(insertOrderMeta, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, 
		order.CustomerID, order.DeliveryService, order.ShardKey, order.SMID, order.DateCreated, order.OOFShard)
		Error(err, "db.Exes order: ")

		_, err = db.Exec(insertDelivery, orDel.Name, orDel.Phone, orDel.Zip, orDel.City, orDel.Address, orDel.Region, orDel.Email)
		Error(err, "db.Exes delivery: ")

		_, err = db.Exec(insertPayment, orPay.Transaction, orPay.RequestID, orPay.Currency, orPay.Provider, orPay.Amount, 
		orPay.PaymentDT, orPay.Bank, orPay.DeliveryCost, orPay.GoodsTotal, orPay.CustomFee) 
		Error(err, "db.Exes payment: ")
			
		_, err = db.Exec(insertItem, orItem.ChrtID, orItem.TrackNumber, orItem.Price, orItem.RID, orItem.Name, orItem.Sale, 
		orItem.Size, orItem.TotalPrice, orItem.NMID, orItem.Brand, orItem.Status)
		Error(err, "db.Exes item: ")

	cache = InsertDataToCache(db)

		}, stan.StartWithLastReceived())

	r := gin.Default()
	r.LoadHTMLFiles("web/userSearch.html")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "userSearch.html", nil)
	})

	r.POST("/search", func(c *gin.Context) {
		searchTerm := c.PostForm("search")

		if data, ok := cache.data[searchTerm]; ok {
			resultJSON, err := json.MarshalIndent(data, "", "\n")
			Error(err, "json.MarshalIndent: ")

			c.HTML(http.StatusOK, "userSearch.html", gin.H{
				"data": string(resultJSON),
			})
			c.JSON(http.StatusOK, gin.H{
				"Данные клиента": cache.data[searchTerm],
			})

		} else {
			c.HTML(http.StatusOK, "userSearch.html", gin.H{})
			 c.String(http.StatusOK, "Данные по указанному ID не найдены")
		}
	})

	r.Run(":8080")
}

func InsertDataToCache(db *sql.DB) *Cache {
	cache := NewCache()
	rows, err := db.Query("SELECT * FROM order_meta")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var orders []Order
	i := 0
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, 
		&order.CustomerID, &order.DeliveryService, &order.ShardKey, &order.SMID, &order.DateCreated, &order.OOFShard); err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
		cache.data[order.OrderUID] = orders[i]
		i++
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return cache
}
