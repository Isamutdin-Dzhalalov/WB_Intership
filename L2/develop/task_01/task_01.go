package main

import (
	"fmt"
	"time"
	"log"
	"github.com/beevik/ntp"
)

func getTime() {

	/* ntp.Time - это функция из библиотеки github.com/beevik/ntp, 
	которая используется для получения текущего времени с NTP сервера.
	"pool.ntp.org" - это адрес NTP сервера.*/

	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}

	/*Format - это метод типа time.Time, который преобразует время в строку 
	в соответствии с указанным форматом.

	time.RFC1123 - константа формата строки из пакета time*/
	fmt.Println(ntpTime.Format(time.RFC1123))
}


func main() {
	getTime()
}
