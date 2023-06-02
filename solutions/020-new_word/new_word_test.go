package newword

import "testing"

func TestNewWord(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "test 1",
			input:    "",
			expected: "Строка пустая",
		},
		{
			name:     "test 2",
			input:    "Liana",
			expected: "iana",
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := newWord(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
