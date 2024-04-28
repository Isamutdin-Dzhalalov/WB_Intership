package main

import (
    "fmt"
)

func changeBit(num int64, i int, newBit int) int64 {
    if newBit == 1 {
        // Установка i-го бита в 1
        return num | (1 << i)
    } else if newBit == 0 {
        // Установка i-го бита в 0
        return num & ^(1 << i)
    } else {
        fmt.Println("Новый бит может быть только 1 или 0")
        return num
    }
}

func main() {
    var num int64 = 22 // Пример числа
    i := 2           // Позиция бита
    newBit := 0      // Значение нового бита (1 или 0)

    newNum := changeBit(num, i, newBit)
    fmt.Println(newNum) // Вывод результата
}

