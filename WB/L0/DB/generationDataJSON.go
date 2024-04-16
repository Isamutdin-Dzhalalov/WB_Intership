package DB

import (
	"log"
    "time"
	"encoding/json"
    "github.com/brianvoe/gofakeit/v7"
)
	

type Order struct {
	OrderUID       string      `json:"order_uid" fake:"{uuid}"`
	TrackNumber    string      `json:"track_number"`
	Entry          string      `json:"entry"`
	Delivery       Delivery    `json:"delivery"`
	Payment        Payment     `json:"payment"`
	Items          [1]Item      `json:"items"`
	Locale         string      `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerID     string      `json:"customer_id"`
	DeliveryService string      `json:"delivery_service"`
	ShardKey       string      `json:"shardkey"`
	SMID           int         `json:"sm_id"`
	DateCreated     string      `json:"date_created"`
	OOFShard       string      `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction   string `json:"transaction"`
	RequestID     string `json:"request_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
	Amount        int    `json:"amount"`
	PaymentDT     int64 `json:"payment_dt"`
	Bank          string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal    int    `json:"goods_total"`
	CustomFee     int    `json:"custom_fee"`
}

type Item struct {
	ChrtID        int    `json:"chrt_id"`
	TrackNumber   string `json:"track_number"`
	Price         int    `json:"price"`
	RID           string `json:"rid"`
	Name          string `json:"name"`
	Sale          int    `json:"sale"`
	Size          string `json:"size"`
	TotalPrice    int    `json:"total_price"`
	NMID          int    `json:"nm_id"`
	Brand         string `json:"brand"`
	Status        int    `json:"status"`
}

func DataGeneration() []uint8 {
    gofakeit.Seed(time.Now().UnixNano())

    var order Order

    err := gofakeit.Struct(&order)
    if err != nil {
		log.Fatal(err)
    }

    jsonData, err := json.MarshalIndent(order, "", " ")
    if err != nil {
		log.Fatal(err)
    }

	return jsonData
}

