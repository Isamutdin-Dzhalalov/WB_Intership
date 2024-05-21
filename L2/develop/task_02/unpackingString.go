package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

var IncorrectString = errors.New("Error: incorrect string")

func main() {
	str := "a4bc2d5e"

	outStr, err := UnpackingString(str)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(outStr)

}

func UnpackingString(inputStr string) (string, error) {

	outStr := make([]byte, 0, len(inputStr))
	// Для хранения последнего символа входной строки.
	lastRune := ""
	for _, v := range inputStr {
		// Преобразуем руну с строку.
		str := string(v)
		// Преобразуем строку в число.
		strRepeats, err := strconv.Atoi(str) 
		if err == nil {
			// Если первый символ - число, возвращаем ошибку.
			if lastRune == "" {
				return "", IncorrectString
			}
			// Добавляем символы нужное количество раз.
			for i := 0; i < strRepeats-1; i++ {
				outStr = append(outStr, lastRune...)
			}

			continue
		}
		
		lastRune = str
		// Дойдя до этого места, символы != числам - добавляем из в результат.
		outStr = append(outStr, lastRune...)
	}

	return string(outStr), nil
}
