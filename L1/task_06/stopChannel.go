/* 
	Остановка выполнения горутин, используя канал, сигнализирущий об остановке.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func printOut(input <-chan int, quit <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		/* Оператор select блокирует выполнение, пока не станет доступна
		хотя бы одна из операций. */
		select {

		// Выполняется условие если данные поступают из канала input.
		case q, ok := <-input:

		/* Проверяем, был ли канал закрыт и если ok == false, значит 
		   канал закрыт и горутина завершает работу. */
			if !ok {
				return
			}
			time.Sleep(200 * time.Millisecond)
			fmt.Println(q)

		// Если канал quit получает значение, горутина завершает работу.
		case <-quit:
			fmt.Println("Канал закрыт")
			return
		}
	}
}

func main() {
	// Создаём экземляр.
	var wg sync.WaitGroup

	input := make(chan int)

	/* Создаём отдельный канал, который будет сигнализировать 
	об остановке горутин, если его значение будет отправлено в канал.*/
	quit := make(chan bool)

	wg.Add(1)
	//Вызываем ф-цию в горутине.
	go printOut(input, quit, &wg)
	
	for i := 0; i < 8; i++ {

	/* Если определённое условие отрабатывает, отправляет значение
		true в канал, сигнализирующий горутинe о неоходимости 
		завершения работы и выходим из цикла. */
		if i == 5 {
			quit <-true
			break
		}
	// Отправляем данные в канал.
		input <- i
	}
	
	wg.Wait()
}
