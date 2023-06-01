package getmiddleelement

import "testing"

func TestGetmiddle(t *testing.T) {
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
			expected: 3,
		},
		{
			name:     "Много чисел",
			input:    []int{3, 6, 23, 78},
			expected: 23,
		},
		{
			name:     "1 число",
			input:    []int{3},
			expected: 3,
		},
		{
			name:     "Отрицательные числа",
			input:    []int{1, -2, -30},
			expected: -2,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := get_middle_element(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
