package addition

import "testing"

func TestAddition(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "test1",
			input:    5,
			expected: 6,
		},
		{
			name:     "test2",
			input:    -3,
			expected: -2,
		},
		{
			name:     "test3",
			input:    1000000,
			expected: 1000001,
		},
		{
			name:     "test4",
			input:    -1000000,
			expected: -999999,
		},
	}

	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := addition(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}

}
