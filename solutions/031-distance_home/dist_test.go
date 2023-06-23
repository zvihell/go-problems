package distancehome

import "testing"

func TestDist(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test 1",
			input:    []int{2, 4, 2, 5},
			expected: 13,
		},
		{
			name:     "test 2",
			input:    []int{-1, -4, -3, -2},
			expected: 10,
		},
		{
			name:     "test 3",
			input:    []int{3, 4, -5, -2},
			expected: 0,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := distance_home(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
