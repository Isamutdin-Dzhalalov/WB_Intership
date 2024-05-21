package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	// Создаем мапу для хранения множеств анаграмм.
	anagramSets := make(map[string][]string)
	hashUniq := make(map[string]bool)
	// Игнорируем повторяющиеся эл-ты с помощью мапы.
	for _, word := range words {
		hashUniq[word] = true
	}
	// Иттерируемся по ключам мапы.
	for word := range hashUniq {
		// Приводим слово к нижнему регистру и сортируем буквы.
		sortedWord := sortLetters(strings.ToLower(word))

		// Если множество для данного отсортированного слова еще не создано, создаем его.
		if _, exists := anagramSets[sortedWord]; !exists {
			anagramSets[sortedWord] = []string{}
		}

		// Добавляем слово в множество.
		anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
	}

	// Удаляем множества, состоящие из одного элемента.
	for key, value := range anagramSets {
		if len(value) == 1 {
			delete(anagramSets, key)
		}
	}

	return anagramSets
}

// Функция для сортировки букв в слове.
func sortLetters(word string) string {
	letters := strings.Split(word, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func main() {
	words := []string{"дaдa", "пятак", "пятка", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramSets := findAnagrams(words)

	for key, value := range anagramSets {
		fmt.Printf("%s: %v\n", key, value)
	}
}
