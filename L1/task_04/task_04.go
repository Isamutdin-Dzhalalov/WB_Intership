package main

import (
	"fmt"
	"time"
//	"sync"
	"log"

	"os"
	"os/signal"
//	"syscall"
)

func workPrint(indx int, dataChan <-chan string) {
	// Читаем данные из канала.
	for data := range dataChan {
		fmt.Printf("Worker: %d, %s\n", indx, data)
	}
}

func main() {
	var numWorkers int
	fmt.Println("Enter the number of workers:")

	
	// Принимаем на ввод число, записываем в переменную
	// и в случае обнаружения ошибок, обрабатываем их.
	if _, err := fmt.Scanf("%d", &numWorkers); err != nil {
		log.Fatal(err)
	} else if numWorkers < 1 {
		log.Fatal("Введите число больше 0")
	}
		

	// Создаём канал.
	dataChan := make(chan string)
	
	for i := 1; i <= numWorkers; i++ {
		// Запускаем функцию в горутине.
		go workPrint(i, dataChan)
	}
	
	/* Создаём канал, который должен быть буферизованным, чтобы избежать пропуска сигнала,
	   если программа не готова его принять в момент отправки сигнала. */
	signals := make(chan os.Signal, 1)

	/* Подписываемся на сигналы.
	   Первым аргументом указываем канал и далее указываем на какие сигналы мы подписываемся. */
	signal.Notify(signals, os.Interrupt)

	go func() {
		<- signals
		close(dataChan)
		log.Fatal("Exit work")
	}()

	for {
		select {
		case <- signals:
			return
		default:
			dataChan <- fmt.Sprintf("Data: %d", time.Now().UnixNano()) // Запись данных в канал
			time.Sleep(time.Second) // Имитация задержки между записями.
		}
	}
}

