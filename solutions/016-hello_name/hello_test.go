package helloname

import "testing"

func TestHello(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "test 1",
			input:    "Нео",
			expected: "Привет, Нео",
		},
		{
			name:     "test 2",
			input:    "Лиана",
			expected: "Привет, Лиана",
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := hello_name(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
