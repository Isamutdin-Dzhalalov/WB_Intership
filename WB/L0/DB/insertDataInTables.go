package DB

import (
//	"database/sql"
//	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)

const (
	setOrderMeta = `
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

	file, err := ioutil.ReadFile("model.json")
	if err != nil {
		log.Fatal("InsertDataInTable -> ioutil.ReadFile: ", err)
	}
	var order Order
	err = json.Unmarshal([]byte(file), &order)
	if err != nil {
		log.Fatal("InsertDataInTable -> json.Unmarshal_1: ", err)
	}

	var delivery Delivery
	err = json.Unmarshal([]byte(file), &delivery)
	if err != nil {
		log.Fatal("InsertDataInTable -> json.Unmarshal_2: ", err)
	}

	var payment Payment
	err = json.Unmarshal([]byte(file), &payment)
	if err != nil {
		log.Fatal("InsertDataInTable -> json.Unmarshal_3: ", err)
	}

	var item Item
	err = json.Unmarshal([]byte(file), &item)
	if err != nil {
		log.Fatal("InsertDataInTable -> json.Unmarshal_4: ", err)
	}

	db, err := ConnectDB()
	if err != nil {
		log.Fatal("InsertDataInTable -> ConnectWithDb: ", err)
	}

	orDel := order.Delivery
	orPay := order.Payment
	orItem := order.Items[0]

	_, err = db.Exec(setOrderMeta, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SMID, order.DateCreated, order.OOFShard)
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
		
}
