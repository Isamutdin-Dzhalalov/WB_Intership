package DB

import (
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	stan "github.com/nats-io/stan.go"

	"html/template"
	"net/http"
//	"main/web"
)

const (
	insertOrderMeta = `
		INSERT INTO order_meta (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

	insertDelivery = `
		INSERT INTO delivery (name, phone, zip, city, address, region, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7);`

	insertPayment = `
		INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`

	insertItem = `
		INSERT INTO item (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

)

func InsertDataInTable() {

	sc, err := stan.Connect("test-cluster", "test-client")
	if err != nil {
		log.Fatal("stan.Connect: ", err)
	}
	defer sc.Close()

	file, err := ioutil.ReadFile("model.json")
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	err = sc.Publish("foo", file)
	if err != nil {
		log.Fatal("sc.Publish: ", err)
	}

	sub, err := sc.Subscribe("foo", func(msg *stan.Msg) {

		log.Println("-------------------------")
		log.Println("Connection NATS-streaming")
		log.Println("-------------------------")

		var order Order
		err = json.Unmarshal([]byte(file), &order)
		if err != nil {
			log.Fatal("InsertDataInTable -> json.Unmarshal_1: ", err)
		}

		db, err := ConnectDB()
		if err != nil {
			log.Fatal("InsertDataInTable -> ConnectWithDb: ", err)
		}

		orDel := order.Delivery
		orPay := order.Payment
		orItem := order.Items[0]

		_, err = db.Exec(insertOrderMeta, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SMID, order.DateCreated, order.OOFShard)
		if err != nil {
			log.Fatal("InsertDataInTable -> db.Exec: ", err)
		}

		_, err = db.Exec(insertDelivery, orDel.Name, orDel.Phone, orDel.Zip, orDel.City, orDel.Address, orDel.Region, orDel.Email)
		if err != nil {
			log.Fatal("InsertDataInTable -> db.Exec: ", err)
		}

		_, err = db.Exec(insertPayment, orPay.Transaction, orPay.RequestID, orPay.Currency, orPay.Provider, orPay.Amount, orPay.PaymentDT, orPay.Bank, orPay.DeliveryCost, orPay.GoodsTotal, orPay.CustomFee) 
		if err != nil {
			log.Fatal("InsertDataInTable -> db.Exec payment: ", err)
		}
			
		_, err = db.Exec(insertItem, orItem.ChrtID, orItem.TrackNumber, orItem.Price, orItem.RID, orItem.Name, orItem.Sale, orItem.Size, orItem.TotalPrice, orItem.NMID, orItem.Brand, orItem.Status)
		if err != nil {
			log.Fatal("InsertDataInTable -> db.Exec payment: ", err)
		}


 	http.HandleFunc("/your_server_script.php", handleForm)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Создаем мапу ключ-значение, где значения являются интерфейсами
		cache := CacheData(order)
		/*
        data := map[string]interface{}{
            "Ключ1": "Значение1",
            "Ключ2": 123,
            "Ключ3": struct{ Name string }{Name: "Значение3"},
        }
		*/

        // Загружаем и разбираем шаблон
    //    tmpl, err := template.ParseFiles("web/template.html")
         tmpl, err := template.ParseFiles("web/userSearch.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Выполняем шаблон, передавая мапу в качестве данных
        err = tmpl.Execute(w, *cache)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })



	}, stan.StartWithLastReceived())
	sub.Unsubscribe()
	//select{}	
    http.ListenAndServe(":8080", nil)
}

func handleForm(w http.ResponseWriter, r *http.Request) {
    // Проверяем, что метод запроса является POST
    if r.Method != http.MethodPost {
        http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        return
    }

    // Парсим данные формы
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Ошибка при парсинге формы", http.StatusBadRequest)
        return
    }

    // Получаем данные из формы
    search := r.FormValue("search")

    fmt.Printf("Вы искали: %s\n", search)
    // Отправляем данные обратно клиенту
    fmt.Fprintf(w, "Вы искали: %s", search)
}
