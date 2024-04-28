package main

import "fmt"

func main() {
	intersection := make([]string, 0, 4)
    // Исходная последовательность строк
    sequence := []string{"cat", "cat", "dog", "cat", "tree"}

    /* создаём мапу, элементы слайса в качестве ключей и устанавливаем
	значение true, тем самым получаем уникальные значения в виде ключей мапы */
    set := make(map[string]bool)
    for _, item := range sequence {
        set[item] = true
    }
	// Достаем ключи из мапы и добавляем в слайс.	
	for value := range set {
		intersection = append(intersection, value)	
	}

	fmt.Println(intersection)
}
