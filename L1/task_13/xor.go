/* Результат выполнения операции XOR (исключающее ИЛИ) равен 1,
   когда один из битов b или a равен 1. В остальных ситуациях 
   результат применения оператора XOR равен 0.*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
    a := 10
    b := 20

	/* Функция FormatInt преобразовывает число с строку
	   указав основание системы счисления(2 - двоичная). */
	binaryA := strconv.FormatInt(int64(a), 2)
	binaryB := strconv.FormatInt(int64(b), 2)

	// Выводим числа в двоичной системе.
	fmt.Println("a: ", binaryA)
	fmt.Println("b: ", binaryB)

    fmt.Printf("До обмена: a: %d, b: %d\n", a, b)

    // Меняем местами a и b операцией XOR.
    a = a ^ b
	/* a =  1010
	   b = 10100
	    == 11110
	Аналогично с другими операциями.
	*/
    b = a ^ b
    a = a ^ b

    fmt.Printf("После обмена: a: %d, b: %d\n", a, b)
}

