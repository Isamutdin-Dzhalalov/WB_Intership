package main

import (
    "fmt"
    "os"
    "sync"
    "time"
	"strconv"
	"log"
)

func worker(id int, wg *sync.WaitGroup, dataChan <-chan string) {
    defer wg.Done()
    for data := range dataChan {
        fmt.Printf("Воркер %d: %s\n", id, data)
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Укажите количество воркеров")
        return
    }

    numWorkers, err := strconv.Atoi(os.Args[1])
    if err != nil {
        log.Println("Неверное количество воркеров")
        return
    }

    dataChan := make(chan string)
    var wg sync.WaitGroup

    // Запуск воркеров
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go worker(i, &wg, dataChan)
    }

    // Постоянная запись данных в канал
    go func() {
        for {
            dataChan <- fmt.Sprintf("Данные %d", time.Now().UnixNano())
            time.Sleep(time.Second) // имитация задержки между записями
        }
    }()

    // Ожидание завершения всех воркеров
    wg.Wait()
}

