package main

import (
	"fmt"
)

/* Ф-ция принимает аргумент типа interface, что позволит
	принимать значения любого типа.*/
func determineType(i interface{}) {
// i.(type) чтобы определить тип переданного значения.
	switch v := i.(type) {
	case int:
		fmt.Printf("Тип: int, значение: %v\n", v)
	case string:
		fmt.Printf("Тип: string, значение: %v\n", v)
	case bool:
		fmt.Printf("Тип: bool, значение: %v\n", v)
	case chan int:
		fmt.Printf("Тип: chan int\n")
	default:
		fmt.Printf("Неизвестный тип: %T\n", i)
	}
}

/*
	Вариант 2:
func getType(v interface{}) string {
//	fmt.Sprintf с форматным спецификатором %T, 
//	который возвращает строку, содержащую тип переданного значения. 
	return fmt.Sprintf("%T", v)
}
*/

func main() {
	var myVar interface{}

	// Примеры значений разных типов
	myVar = 10
	determineType(myVar)

	myVar = "Hello, World!"
	determineType(myVar)

	myVar = true
	determineType(myVar)

	// Пример с каналом
	ch := make(chan int)
	myVar = ch
	determineType(myVar)
}
