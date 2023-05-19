package multiplymaxmin

import (
	"fmt"
	"testing"
)

func Test_multiplyMaxMin(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Пустой массив",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Простой",
			input:    []int{2, 3},
			expected: 6,
		},
		{
			name:     "Много чисел",
			input:    []int{3, 3, 40, 4, 5, 20},
			expected: 120,
		},
		{
			name:     "1 число",
			input:    []int{3},
			expected: 9,
		},
		{
			name:     "Отрицательные числа",
			input:    []int{1, -2, -30},
			expected: -30,
		},
	}

	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := multiplyMaxMin(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			fmt.Printf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}

}
