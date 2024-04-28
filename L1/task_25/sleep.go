package main

import (
	"fmt"
	"time"
)

/*
   time.After(d) создает канал типа <-chan time.Time, который 
   будет отправлять значение типа time.Time после истечения времени d.
   <-time.After(d) внутри функции sleep приостанавливает выполнение программы 
   на указанное время d, пока канал не будет готов (то есть не отправит значение).
*/

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Начало паузы")
	sleep(2 * time.Second) 
	fmt.Println("Конец паузы")
}

