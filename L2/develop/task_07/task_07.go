package main

import (
	"fmt"
	"sync"
	"time"
)

var stdoutMutex = &sync.Mutex{}

// Ф-ция принимает канал интерфейсов для чтения).
func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	// Для ожидания завершения горутин.
	var wg sync.WaitGroup

	for _, c := range channels {
		start := time.Now()
		wg.Add(1)
		go func(c <-chan interface{}) {
			for v := range c {
				out <- v
			}
			/* В Go, стандартный вывод(fmt.Println()) являются потоками,
			которые предназначены для использования только одной горутиной за раз.
			Попытка записать в stdout из другой горутины может привести к блокировке.
			Чтобы избежать конфликтов при одновременном обращении из разных горутин,
			используем Mutex для синхронизации доступа.*/

			stdoutMutex.Lock()
			fmt.Printf("close channel after %v\n", time.Since(start))
			stdoutMutex.Unlock()
			wg.Done()
		}(c)
	}

	wg.Wait()
	close(out)

	return out
}
/* Ф-ция создаёт канал, запускает горутину которая ждёт after(время).
   После завершения канал закрывается и возвращает канал ans.*/
func sig(after time.Duration) <-chan interface{} {
	ans := make(chan interface{})
	go func() {
		defer close(ans)
		time.Sleep(after)
	}()

	return ans
}


func main() {
	start := time.Now()
	// Вызываем функцию or  с 5-ю аргументами(рез-ми функции sig).
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(7*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}
