package google

import "testing"

func TestGoogle(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "test1",
			input:    5,
			expected: "Gooooogle",
		},
		{
			name:     "test2",
			input:    0,
			expected: "Ggle",
		},
		{
			name:     "test3",
			input:    2,
			expected: "Google",
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := google(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}

}
