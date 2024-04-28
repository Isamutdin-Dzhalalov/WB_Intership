package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
	"log"
)

func main() {
	// Передаём желаемое время в аргументах.
	args := os.Args
	// Проверяем количество введённых аргументов.
	if len(args) < 2 {
		log.Fatal("Введите время")
	}

	chanUser := make(chan int)
	go writeChan(chanUser) 
	go readChan(chanUser)
	defer close(chanUser)

	// Преобразовываем второй переданый аргумент в программу из строки в целое число.
	seconds, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}
	// Создаём таймер для завершения программы.
	timeChan := time.NewTimer(time.Duration(seconds) * time.Second)
	// Таймер не даёт завершиться программе на заданное время.
	<- timeChan.C
	fmt.Println("Время истекло!")
	return
}

func readChan(chanUser <-chan int) {
	for out := range chanUser {
//	for {
/*
		out, ok := <-chanUser
		// Если канал закрыт - выходим из цикла.
		if !ok {
			break
		}
*/
		fmt.Println(out)
	}
}

func writeChan(chanUser chan<- int) {
	for {
		timeNow := int(time.Now().UnixNano() / int64(time.Second))
		// Отправляем данные в канал.
		chanUser <- timeNow
		time.Sleep(1 * time.Second)
	}
}
	
