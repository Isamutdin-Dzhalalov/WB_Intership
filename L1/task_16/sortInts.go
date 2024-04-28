package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 2, 9, 1, 7, 3, 8, 4, 6}
	fmt.Println("Unsorted array:", numbers)

	sort.Ints(numbers)

	fmt.Println("Sorted array:", numbers)
}
