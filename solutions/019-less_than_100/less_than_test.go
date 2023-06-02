package lessthan100

import "testing"

func TestLess(t *testing.T) {
	testCases := []struct {
		name     string
		input1   int
		input2   int
		expected bool
	}{
		{
			name:     "test1",
			input1:   22,
			input2:   15,
			expected: true,
		},
		{
			name:     "test2",
			input1:   83,
			input2:   34,
			expected: false,
		},
		{
			name:     "test3",
			input1:   3,
			input2:   77,
			expected: true,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := less_than_100(tc.input1, tc.input2)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}

}
