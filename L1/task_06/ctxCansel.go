package main

import (
	"time"
	"context"
	"log"
	"sync"
)

func main() {
	//Создаём контекст с функцией отмены.
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	
	// Запускаем горутину.
	go worker(ctx, &wg)

	//Ждём некоторое время, затем отменяем контекст вызывая cancel().
	time.Sleep(3 * time.Second)
	//Вызываем cancel(), чтобы отменить контекст и завершить горутину.
	cancel()
	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		// Горутина завершается, после того, как контекст был отменён.
		case <-ctx.Done():
			log.Println("Работа горутины завершена")
			return
		default:
			log.Println("Горутина в работе")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
			
	
