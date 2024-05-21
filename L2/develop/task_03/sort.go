package main

import (
	"bufio"
	"io"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Arguments struct {
	k int
	n bool
	r bool
	u bool
	M bool
	b bool
	c bool
	h bool
}

var err error

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите название файла:")
	var input string
	if scanner.Scan() {
		input = scanner.Text()
	} else {
		log.Fatal("Некорректные данные")
	}

	var Arg Arguments
	// Слайс для хранения строк из файла.
	lines := make([]string, 0)

	flag.IntVar(&Arg.k, "k", 0, "Сортировка по определенному столбцу")
	flag.BoolVar(&Arg.n, "n", false, "Сортировка по числовому значению")
	flag.BoolVar(&Arg.r, "r", false, "Сортировка в обратном порядке")
	flag.BoolVar(&Arg.u, "u", false, "Не выводить повторяющиеся строки")
	flag.BoolVar(&Arg.M, "M", false, "Сортировка по названию месяца")
	flag.BoolVar(&Arg.b, "b", false, "Игнор хвостовых пробелов")
	flag.BoolVar(&Arg.c, "c", false, "Проверка сортировки")
	flag.BoolVar(&Arg.h, "h", false, "Сортировка по числовому значению с учетом суффиксов")
	// Заполняем поля структуры.
	flag.Parse()

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	lines, err = ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	result := FullSort(lines, Arg)
//	fmt.Println(result)
	recordFile("outText.txt", result)
}

/* Вызываем ф-ции в зависимости от найденных флагов.
   Если флаги не найденны, сортируем а алфавитном порядке.*/
func FullSort(lines []string, Arg Arguments) []string {
	var result []string
	switch {
	case Arg.k > 0:
		result = SortByColumn(Arg.k, lines)
	case Arg.n:
		result, err = NumberSort(lines)
		if err != nil {
			fmt.Println(err)
		}
	case Arg.r:
		result = ReverseSort(lines)
	case Arg.u:
		result = IgnoreDublicateLines(lines)
	case Arg.M:
		result = MonthSort(lines)
	case Arg.c:
		Check := CheckSort(lines)
		if Check {
			fmt.Println("Данные отсортированы")
		} else {
			fmt.Println("Данные не отсортированы")
		}
	case Arg.h:
		result = NumberSortWithSuffix(lines)
	default:
		sort.Strings(lines)
		result = lines
	}

	return result
}

/* Ф-ция для записи эл-тов массива в файл.
   os.Create - создаёт файл или перезаписывает имеющийся.
   перед выходом из ф-ции закрываем файл "defer..."*/

func recordFile(file string, array []string) {
	outFile, err := os.Create(file)
	if err != nil {
		log.Fatal("recordFile", err)
	}
	defer outFile.Close()

	for i := 0; i < len(array)-1; i++ {
		outFile.WriteString(array[i] + "\n")
	}
	// Последний эл-нт массива записываем без переноса строки '\n'.
	outFile.WriteString(array[len(array)-1])
}

//Чтение файла и получение строк
func ReadFile(file io.Reader) ([]string, error) {
	var lines []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		lines = append(lines, line)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return lines, nil
}

/* Удаление дублированных строк.
   С помощью структуры map добиваемся получения
   уникальных строк.*/
func IgnoreDublicateLines(lines []string) []string {
	hash := make(map[string]bool)
	for _, line := range lines {
		hash[line] = true
	}
	result := make([]string, 0, len(hash))
	for key, _ := range hash {
		result = append(result, key)
	}

	return result
}

//Разбиение строки на отдельные слова
func WordsFromString(numColumn int, line string) string {
	if line == "" {
		return ""
	}
	Words := strings.Fields(line)
	ln := utf8.RuneCountInString(line)
	if numColumn < ln {
		return Words[numColumn]
	} else {
		return ""
	}

}

//Сортировка по отдельной колонке
func SortByColumn(num int, lines []string) []string {
	ReceivedWords := make([]string, 0)
	Result := make([]string, 0)
	for _, line := range lines {
		//Получение слова из определенного столбца
		TheRightWord := WordsFromString(num, line)
		ReceivedWords = append(ReceivedWords, TheRightWord)
	}
	sort.Strings(ReceivedWords)
	for _, w := range ReceivedWords {
		for _, line := range lines {
			if strings.Contains(line, w) {
				Result = append(Result, line)
			}
		}
	}
	return Result
}

// Сортировка по месяцам.
func MonthSort(lines []string) []string {
	var months = map[string]int{
		"JAN": 1,
		"FAB": 2,
		"MAR": 3,
		"APR": 4,
		"MAY": 5,
		"JUN": 6,
		"JUL": 7,
		"AUG": 8,
		"SEP": 9,
		"OCT": 10,
		"NOV": 11,
		"DEC": 12,
	}
	notMoth := make([]string, 0)
	answerInt := make([]int, 0)
	result := make([]string, 0)
	for _, month := range lines {
		if val, ok := months[month]; ok {
			answerInt = append(answerInt, val)
		} else {
			notMoth = append(notMoth, month)
		}

	}
	sort.Ints(answerInt)
	for _, num := range answerInt {
		for key, _ := range months {
			if months[key] == num {
				result = append(result, key)
			}
		}
	}
	for _, v := range notMoth {
		result = append(result, v)
	}
	return result
}

// Ф-ция для cортировки по числам.
func NumberSort(lines []string) ([]string, error) {
	nums := make([]int, 0)
	result := make([]string, 0)
	var n int
	for _, num := range lines {
		for _, symbol := range num {
		/* Игнорируем пробел, далее проверяем, является ли символ числовым,
		   если да, добавляем слайс int`ов, 
		   если нет, строку добавляем в слайс string и выходим из цикла.*/
			if symbol == ' ' {
			} else if rune(symbol) >= '0' && rune(symbol) <= '9' { 
				n, _ = strconv.Atoi(string(symbol))
				nums = append(nums, n)
			} else {
				result = append(result, num)
				break
			}
		}
	}
	// Сортируем слайс чисел и добавляем в итоговый слайс.
	sort.Ints(nums)
	for _, val := range nums {
		ns := strconv.Itoa(val)
		result = append(result, ns)
	}
	return result, nil
}

// Сортировка по числовому значению с учетом суффиксов.
func NumberSortWithSuffix(Lines []string) []string {
	sort.Strings(Lines)
	return Lines
}

// Проверка сортировки.
func CheckSort(lines []string) bool {
	answer := sort.StringsAreSorted(lines)
	return answer
}

// Сортировка в обратном порядке.
func ReverseSort(lines []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(lines)))
	return lines
}

