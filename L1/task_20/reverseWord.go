package main

import (
	"fmt"
	"strings"
)

/*
	Используем функцию strings.Fields() для разбиения строки на срез слов, разделенными пробелами.
	Создаём два указателя(начало, конец), итерируяс в цикле, меняем слова местами.
	Возвращаем строку в вызываещую ф-цию, объединив срез слов в строку, ф-цией strings.Join()
	где второй аргумент - разделитель слов.
*/

func swapWords(str string) string {
	words := strings.Fields(str)
	l, r := 0, len(words)-1
	for l < r { 
		words[l], words[r] = words[r], words[l]
		l++
		r--
	}
	return strings.Join(words, " ")
}

func main() {
	str := "отрабатывает) вроде Реализация"
	str = swapWords(str)

	fmt.Println(str)
}

