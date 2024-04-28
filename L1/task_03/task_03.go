package main

import (
	"fmt"
	"sync"
)

func sumNumbers(numbers []int) int {
	
	var wg sync.WaitGroup
	sumOfSquare := make(chan int)
	wg.Add(len(numbers))

	for _, num := range numbers {
		go func(n int) {
			defer wg.Done()
			sumOfSquare <- n * n
		}(num)
	}

	go func() {
		wg.Wait()
		close(sumOfSquare)
	}()
	
	var result int
	for num := range sumOfSquare {
		result += num
	}

	return result
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	result := sumNumbers(numbers)
	fmt.Println(result)

}


