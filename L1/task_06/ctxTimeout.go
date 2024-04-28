package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Создаем контекст с таймаутом.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for {
		select {
		//Горутина завершает работу, если поступил сингал отмены ctx.Done().
		case <-ctx.Done():
			fmt.Println("Контекст отменен, горутина завершена")
			return
		default:
			fmt.Println("Горутина работает!")
		}
			time.Sleep(200 * time.Millisecond)
	}
}
