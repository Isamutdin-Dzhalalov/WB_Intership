package main

import "fmt"
// Ф-ция принимает слайс и индексы начала и конца подмассива.
func partition(arr []int, low, high int) ([]int, int) {
// Последний эл-нт в качестве опорного.
    pivot := arr[high]
    i := low
    for j := low; j < high; j++ {
        if arr[j] < pivot {
		//Меняем эл-ты местами.
            arr[i], arr[j] = arr[j], arr[i]
            i++
        }
    }
    arr[i], arr[high] = arr[high], arr[i]
    return arr, i
}

// Ф-ция рекурсивно вызывает себя, разбивая массив на 2 части, сортируя их.
func quickSort(arr []int, low, high int) []int {
    if low < high {
        var p int
        arr, p = partition(arr, low, high)
        arr = quickSort(arr, low, p-1)
        arr = quickSort(arr, p+1, high)
    }
    return arr
}

func main() {
    arr := []int{5, 6, 7, 2, 1, 0}
	fmt.Println(arr)
    fmt.Println(quickSort(arr, 0, len(arr) - 1))
}

