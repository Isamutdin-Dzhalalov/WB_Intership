package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	// Тестовые данные для проверки чтения файла.
	testCases := []struct {
		input    string   // Входные данные (строка для чтения).
		expected []string // Ожидаемый результат (срез строк).
	}{
		{
			input:    "line1\nline2\nline3",
			expected: []string{"line1", "line2", "line3"},
		},
	}

	for _, tc := range testCases {
		// Создаем входной поток данных.
		reader := strings.NewReader(tc.input)

		// Вызываем тестируемую функцию.
		result, err := ReadFile(reader)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Сравниваем результат с ожидаемым результатом.
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Unexpected result. Expected: %v, got: %v", tc.expected, result)
		}
	}
}

