package main

import "fmt"

func main() {
	// создаём две переменные
	a := 5
	b := 10
	fmt.Printf("a: %d, b: %d\n", a, b)
	//1. меняем их значения местами
	a, b = b, a
	fmt.Printf("a: %d, b: %d\n", a, b)


	//2. Меняем обратно методом арифметический операций.
	b = a + b
    a = b - a
    b = b - a
	fmt.Printf("a: %d, b: %d\n", a, b)
}
