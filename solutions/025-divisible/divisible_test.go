package divisible

import "testing"

func TestDivisible(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "test 1",
			input:    1,
			expected: false,
		},
		{
			name:     "test 2",
			input:    1000,
			expected: true,
		},
		{
			name:     "test 3",
			input:    100,
			expected: true,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := divisible(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
