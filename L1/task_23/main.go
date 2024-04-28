package main

import (
	"fmt"
	"os"
)

/* 
   Меняем местами значение по индексу с последним значением в слайсе,
   далее присваиваем слайсу, срез слайса без последнего элемента,
   тем самым удаляя последний элемент.
   Временная сложность O(1), но нужно учитывать, что расположение элементов меняется.
*/

func swapElem(arr []int, indx int) []int {
	last := len(arr)-1
	arr[indx] = arr[last]
	arr = arr[:last]
	return arr
}

/* 
   Удаляем элемент по индексу, путём объединения двух частей исходного слайса,
   исключая значение по индексу.
   Временная сложность O(n), при этом расположение элементов не меняется.
*/

func deleteViaAppend(arr []int, indx int) []int {
	arr = append(arr[:indx], arr[indx+1:]...)
	return arr
}
		
// Проверяем, вы выходит ли ввёденный индекс за диапазон массива.
func correctIndx(arr []int, indx int) bool {
	if indx < 0 || indx > len(arr)-1 {
		fmt.Println("Введённый индекс выходит за пределы диапазона массива")
		return false
	}
	return  true
}

func input() int {
	fmt.Printf("Введите индекс: ")
	var indx int
	fmt.Fscan(os.Stdin, &indx)
	return indx
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	indx := input()
//Если ф-ция возвращает false, то останавливаем программу.
	if ok := correctIndx(arr, indx); !ok {
		return
	}
	resSwap := swapElem(arr, indx)
	resDelete := deleteViaAppend(arr, indx)
	fmt.Println("Результат функции resSwap: ", resSwap)
	fmt.Println("Результат функции deleteViaAppend: ", resDelete)
}
