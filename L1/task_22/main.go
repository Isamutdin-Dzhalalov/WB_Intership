package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем большие числа.
	a := big.NewInt(1 << 21) // 2^21
	b := big.NewInt(1 << 21) // 2^21

	// Выполняем арифметические операции.
	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)

	// Выводим результаты.
	fmt.Printf("a = %d\nb = %d\n", a, b)
	fmt.Printf("a + b = %d\n", sum)
	fmt.Printf("a - b = %d\n", diff)
	fmt.Printf("a * b = %d\n", prod)
	fmt.Printf("a / b = %d\n", div)
}
