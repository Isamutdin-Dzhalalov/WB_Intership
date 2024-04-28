package main

import (
	"fmt"
)

func createSlice(lenSlice int) []int {
	sliceNum := make([]int, lenSlice)
	for i := 0; i < lenSlice; i++ {
		sliceNum[i] = i
	}
	return sliceNum
}
// Записываем данные в канал, перед выходом из ф-ции закрываем канал.
func addCh1(ch1 chan<- int, sliceNum []int) {
	defer close(ch1)
	for _, value := range sliceNum {
		ch1 <-value
	}
}
// Записываем в канал значения*2, перед выходом из ф-ции закрываем канал.
func multiplication(ch2 chan<- int, sliceNum []int) {
	defer close(ch2)
	for _, value := range sliceNum {
		ch2 <-value * value
	}
}

func main() {
	lenSlice := 7
	//Вызываем ф-цию для создания и заполнения слайса.
	sliceNum := createSlice(lenSlice)
	// Создаём два канала.
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Вызываем ф-ции в горутине.
	go addCh1(ch1, sliceNum)
	go multiplication(ch2, sliceNum)

	//Читаем данные из канала.
	for value := range ch2 {
		fmt.Println(value)
	}
}
