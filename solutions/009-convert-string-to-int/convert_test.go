package convertstring

import "testing"

func TestConvertString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "test1",
			input:    "1",
			expected: 1,
		},
		{
			name:     "test2",
			input:    "10",
			expected: 10,
		},
	}

	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := convert_string_2_int(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
