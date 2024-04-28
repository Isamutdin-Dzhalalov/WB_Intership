package main

import (
	"fmt"
	"strings"
)

/*
  Ф-ция преобразует строку в нижний регистр strings.ToLower(), далее
  в цикле делаем проверку на наличия значения(true) в мапе по ключу, 
  где ключ - символ строки. Если ok == true, значит символ ранее нам 
  встречался и мы выходим из ф-ции и возвращаем false, иначе возвращаем true.
*/

func uniq(str string) bool {
	str = strings.ToLower(str)
	uniqHash := make(map[rune]bool)

	for _, value := range str {
		if _, ok := uniqHash[value]; ok {
			return false
		}
		uniqHash[value] = true
	}
	return true
}

func main() {
	str := "Abcdqksv"
	fmt.Println(uniq(str))
}
