package convertminutes

import "testing"

func TestConvertMinutes(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "test1",
			input:    1,
			expected: 60,
		},
		{
			name:     "test2",
			input:    2,
			expected: 120,
		},
		{
			name:     "test3",
			input:    0,
			expected: 0,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := convert_minutes_2_seconds(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
