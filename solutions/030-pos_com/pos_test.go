package poscom

import "testing"

func TestPos(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected float64
	}{
		{
			name:     "test1",
			input:    3,
			expected: 8,
		},
		{
			name:     "test2",
			input:    1,
			expected: 2,
		},
		{
			name:     "test3",
			input:    10,
			expected: 1024,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := pos_com(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
