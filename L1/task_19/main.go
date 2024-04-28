package main

import (
	"fmt"
)

/* В Go нельзя напрямую брать адрес элемента массива или срезa, если это значение типа "byte".
	Строка и является последовательностью байт, а не сивмолов. Символ может занимать более 1 байта.
	Поэтому преобразуем строку в срез рун([]rune). 

	Затем проинициализируем два указателя, указывающие на первый и последний символ.
	Далее, в цикле, будем менять символы местами до тех пора, пока указатели не "столкнуться"
	и вернём строку в вызывающую ф-цию.
*/

func reverseString(str string) string {
	revers := []rune(str)
	l, r := 0, len(revers) - 1
	for l < r {
		revers[l], revers[r] = revers[r], revers[l]
		l++
		r--
	}
	return string(revers)
}

func main() {
	str := "Привет"
	
	result := reverseString(str)
	fmt.Println(result)
}

/*

package main

import (
    "fmt"
    "github.com/rivo/uniseg"
)

func main() {
    str := "главрыба — абырвалг"
    reversed := uniseg.ReverseString(str)
    fmt.Println(reversed)
}

