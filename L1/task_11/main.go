package main

import (
	"fmt"
)

func intersection(slice1, slice2 []int) []int {
	var resultSlice []int 
	m := make(map[int]bool)
	/* Ключ в мапе - это значения первого множества,
		которыe помечаем, как true. */
	for _, value := range slice1 {
		m[value] = true
	}
	// Проходимся циклом по второму множеству.
	for _, value := range slice2 {
	/* Проверяем условием, если значение по ключу == true,
	   значит такое же значение есть в первом множестве.
	   Добавляем это значение мн-ва в слайс. */
		if m[value] {
			resultSlice = append(resultSlice, value)
		}
	}
	//Возвращаем результат в виде слайса.
	return resultSlice
}

func main() {
	
	slice1 := []int{1, 3, 5, 7}
	slice2 := []int{5, 6, 3, 2}
	result := intersection(slice1, slice2)
	fmt.Println(result)
}

	
