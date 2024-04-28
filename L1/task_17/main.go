package main

import "fmt"

// Ф-ция возвращает -1 если число не найдённо, иначе индекс ячейки найденного значения.
func binarySearch(nums []int, target int) int{
	leftPtr := 0
	rightPtr := len(nums) - 1
	// Кладём в переменную значение из середины массива.
	middlePtr := (rightPtr + leftPtr) / 2 
/*
	Если число которое ищем, меньше крайне левого числа или больше крайне правого числа в массиве,
	это говорит о том, что данного числа нет в массиве, так как массив отсортирован.
	Выходим из ф-ции.
*/
	if nums[leftPtr] > target || nums[rightPtr] < target {
		return -1
	}

	/* Выходим из цикла, если указатели "столкнутся".
	   В цикле делаем проверки: если число, которое ищем, больше middlePtr, то отсекаем 
	   левую часть массива включая middlePtr, если число меньше middlePtr, то отсекаем правую часть
	   включая middlePtr, если оба условия не верны, значит middlePtr и есть число, которое ищем.
	*/
	for leftPtr <= rightPtr {
		if nums[middlePtr] > target {
			rightPtr = middlePtr - 1
		} else if nums[middlePtr] < target {
			leftPtr = middlePtr + 1 
		} else {
			return middlePtr
		}
	/* В каждой итерации ищем число, которое находится посередине диапазона массива между
		левым и правым индексами. */
		middlePtr = (rightPtr + leftPtr) / 2 
	}
	return -1
}

func main() {
	nums := []int{-1,0,3,5,9,12}
	target := 12
	res := binarySearch(nums, target)
	if res == -1 {
		fmt.Println("Число не найденно.")
	} else {
		fmt.Println("Число в ячейке с индексом ", res)
	}
}
