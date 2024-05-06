package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mut sync.Mutex
	NumMap map[int]int
}

func (s *SafeCounter) AddNum(i int) {
	/* При вызове Lock(), горутина блокируется, пока
	 не получит эксклюзивный доступ к ресурсу, тем самым
	 предотвращая возникновение гонки*/
	s.mut.Lock()

	/* Перед выходом из функции вызываем Unlock(),
	тем самым разблокировав доступ к данным для 
	других горутин */
	defer s.mut.Unlock()
	s.NumMap[i] = i
}

func main() {
	var wg sync.WaitGroup
	s := SafeCounter{NumMap: make(map[int]int)}
	for i := 0; i < 7; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.AddNum(i)
		}(i)
	}
	wg.Wait()
	fmt.Println(s.NumMap)
}
