package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Определение флагов.
	fields := flag.String("f", "", "Выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "Использовать другой разделитель (по умолчанию табуляция)")
	separated := flag.Bool("s", false, "Только строки с разделителем")
	flag.Parse()

	if *fields == "" {
		fmt.Println("Необходимо указать хотя бы одно поле для выделения с помощью флага -f")
		os.Exit(1)
	}

	// Получение списка колонок.
	selectedFields := parseFields(*fields)

	// Чтение из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Проверка на наличие разделителя в строке.
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		// Разделение строки на колонки.
		columns := strings.Split(line, *delimiter)
		var selectedValues []string

		// Получение значений по указанным колонкам.
		for _, field := range selectedFields {
			if field >= 0 && field < len(columns) {
				selectedValues = append(selectedValues, columns[field])
			}
		}

		// Вывод выбранных колонок.
		fmt.Println(strings.Join(selectedValues, *delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения строки:", err)
	}
}

// parseFields преобразует строку с номерами полей в слайс целых чисел.
func parseFields(fields string) []int {
	fieldList := strings.Split(fields, ",")
	var result []int

	for _, field := range fieldList {
		index, err := strconv.Atoi(field)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка парсинга индекса поля %s: %v\n", field, err)
			os.Exit(1)
		}
		// Преобразуем в нулевой индекс.
		result = append(result, index-1)
	}

	return result
}

