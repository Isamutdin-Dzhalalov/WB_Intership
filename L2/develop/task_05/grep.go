package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	 A int
	 B int
	 C int
	 c bool
	 i bool
	 v bool
	 F bool
	 n bool
)

func main() {
	subStr, fileName := parseFlag()

	res := grep(fileName, subStr)
	for _, str := range res {
		fmt.Println(str)
	}
}

func parseFlag() (string, string) {
	
	flag.IntVar(&A, "A", 0, "after печатать +N строк после совпадения")
	flag.IntVar(&B, "B", 0, "before печатать +N строк до совпадения")
	flag.IntVar(&C, "C", 0, "context (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&c, "c", false, "count (количество строк)")
	flag.BoolVar(&i, "i", false, "ignore-case (игнорировать регистр)")
	flag.BoolVar(&v, "v", false, "invert (вместо совпадения, исключать)")
	flag.BoolVar(&F, "F", false, "fixed, точное совпадение со строкой, не паттерн")
	flag.BoolVar(&n, "n", false, "line num, напечатать номер строки")
	flag.Parse()

	a := os.Args[1:]
	if len(a) == 0 {
		fmt.Println("go run main.go -v [pattern] [file]")
		log.Fatal()
	}

	if len(flag.Args()) < 2 {
		fmt.Println("go run main.go -v [pattern] [file]")
		log.Fatal()
	}

	subStr := flag.Arg(0)
	fileName := flag.Arg(1)

	return subStr, fileName
}

func grep(fileName string, subStr string) []string {
	ignoreCase := i
	file := openReadFile(fileName)
	var result []string

	if ignoreCase {
		subStr = strings.ToLower(subStr)
	}

	if c {
		i := countStr(file, subStr, ignoreCase)
		result = append(result, fmt.Sprintf("count: %d", i))
		return result
	}

	if B > 0 && C == 0 {
		result = beforeN(file, subStr, B, ignoreCase, n)
	}

	if A > 0 && C == 0 {
		result = afterN(file, subStr, A, C, ignoreCase, n)
	}

	if C > 0 {
		A = C
		B = C
		result = beforeN(file, subStr, B, ignoreCase, n)
		result = append(result, afterN(file, subStr, A, C, ignoreCase, n)...)
	}

	if A == 0 && B == 0 {
		result = allStr(file, subStr)
	}
	return result
}

func allStr(file []byte, subStr string) []string {
	ignoreCase := i
	invert := v
	var result []string
	var contains bool
	// Разделение строки на части по символу новой строки.
	splitLine := strings.Split(string(file), "\n")

	for i, v := range splitLine {
	/* Если ignoreCase то строка преобразуется с ниж.регистр
	  Затем функция strings.Contains проверяет, содержит ли 
	  эта строка подстроку subStr.
	  Иначе проверяем содержание строки без преобразование в ниж.регистр.*/
		if ignoreCase {
			contains = strings.Contains(strings.ToLower(v), subStr)
		} else {
			contains = strings.Contains(v, subStr)
		}
	/* Добавляем строку в result в зависимости от условий.
	   Если n == true добавляем номер строки.*/
		if contains && !invert && !F {
			if n {
				result = append(result, fmt.Sprintf("%d: %s", i+1, v))
			} else {
				result = append(result, v)
			}

		} else if invert && !contains {
			if n {
				result = append(result, fmt.Sprintf("%d: %s", i+1, v))
			} else {
				result = append(result, v)
			}
		}
		if F && v == subStr {
			if n {
				result = append(result, fmt.Sprintf("%d: %s", i+1, v))
			} else {
				result = append(result, v)
			}
		}
	}
	// Если по итогу result пустой.
	if len(result) < 1 {
		fmt.Println("no matches found")
		return nil
	}
	return result
}

func afterN(file []byte, subStr string, A int, C int, ignoreCase bool, n bool) []string {
	var after []string
	j := 0
	var index int
	var contains bool

	splitLine := strings.Split(string(file), "\n")

	for i, v := range splitLine {
		if ignoreCase {
			contains = strings.Contains(strings.ToLower(v), subStr)
		} else {
			contains = strings.Contains(v, subStr)
		}
		if contains {
			if j == 0 {
				index = i
				break
			}
		}
	}

	if index > 0 && A <= index {
		if C > 0 {
			after = append(after, splitLine[index+1:index+A+1]...)
		} else {
			after = append(after, splitLine[index:index+A+1]...)
		}
		// Если флаг n добавляем номера строк.
		if n {
			for i := 1; i < len(after)+1; i++ {
				namberLine := index + i
				if C > 0 {
					after[i-1] = fmt.Sprintf("%d: %s", namberLine+1, after[i-1])
				} else {
					after[i-1] = fmt.Sprintf("%d: %s", namberLine, after[i-1])
				}
			}
		}
	}

	// Если строк в файле меньше, чем запрошено.
	if len(after) < A || A == 0 {
		fmt.Println("A: value of requested strings exceeded")
		return nil
	}

	return after
}

func beforeN(file []byte, subStr string, B int, ignoreCase bool, n bool) []string {
	var before []string
	var index int
	j := 0
	var contains bool

	splitLine := strings.Split(string(file), "\n")

	for i, v := range splitLine {

		if ignoreCase {
			contains = strings.Contains(strings.ToLower(v), subStr)
		} else {
			contains = strings.Contains(v, subStr)
		}

		if contains {
			if j == 0 {
				index = i
				break
			}
		}
	}

	if index > 0 && B <= index {
		before = append(before, splitLine[index-B:index+1]...)

		if n {
			for i := 1; i < len(before)+1; i++ {
				namberLine := index - B + i
				before[i-1] = fmt.Sprintf("%d: %s", namberLine, before[i-1])
			}
		}
	}

	if len(before) < B || B == 0 {
		fmt.Println("B: value of requested strings exceeded")
		return nil
	}
	return before
}

func countStr(file []byte, subStr string, ignoreCase bool) int {
	splitLine := strings.Split(string(file), "\n")
	var i int
	for _, v := range splitLine {
		if ignoreCase {
			if strings.Contains(strings.ToLower(v), subStr) {
				i++
			}
		} else {
			if strings.Contains(v, subStr) {
				i++
			}
		}
	}
	return i
}

func openReadFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	readFile, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return readFile
}
