package difference

import "testing"

func TestDifference(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Пустой слайс",
			input:    []int{0},
			expected: 0,
		},
		{
			name:     "Много чисел",
			input:    []int{10, 15, 20, 2, 10, 6},
			expected: 18,
		},
		{
			name:     "Отрицательные числа",
			input:    []int{-3, 4, -9, -1, -2, 15},
			expected: 24,
		},
		{
			name:     "Много чисел",
			input:    []int{4, 17, 12, 2, 10, 2},
			expected: 15,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := difference(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
