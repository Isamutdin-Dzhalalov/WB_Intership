package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	num int
}

/* Метод позволяет инкрементировать переменную,
   блокировать/разблокировать доступ к num.*/
func (c *Counter) increment() {
	c.mu.Lock()
	c.num++
	c.mu.Unlock()
}

func start(count *Counter, completed chan<- int) {
	var wg sync.WaitGroup

	/* Запускаем 100 горутин, в каждой итерации увеличиваем
	   счётчик горутин(wg.Add(1)) и уменьшает(wg.Done()). */
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(count *Counter, wg *sync.WaitGroup) {
			count.increment()
			wg.Done()
		}(count, &wg)
	}

	// Блокируем дальнейшее выполнение, пока все горутины не завершаться.
	wg.Wait()
	completed <- 1 //Отправляет сигнал о завершении работы всех горутин.
}

func (c *Counter) output() {
	fmt.Println(c.num)
}

func main() {
	counter := Counter{num: 0}
	completed := make(chan int)

	// Запускаем горутину.
	go start(&counter, completed)

	<-completed // Блокируемся, пока не придёт сигнал из ф-ции start.
	counter.output() // После получения сигнала, вызываем метод и выводим результат.
}
