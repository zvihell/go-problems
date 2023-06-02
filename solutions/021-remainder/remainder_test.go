package remainder

import "testing"

func TestRemainder(t *testing.T) {
	testCases := []struct {
		name     string
		input1   int
		input2   int
		expected int
	}{
		{
			name:     "test1",
			input1:   1,
			input2:   3,
			expected: 1,
		},
		{
			name:     "test2",
			input1:   5,
			input2:   5,
			expected: 0,
		},
		{
			name:     "test3",
			input1:   7,
			input2:   2,
			expected: 1,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := remainder(tc.input1, tc.input2)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}

}
