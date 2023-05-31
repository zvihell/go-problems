package converthours

import "testing"

func TestConvertHours(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "test1",
			input:    2,
			expected: 7200,
		},
		{
			name:     "test2",
			input:    24,
			expected: 86400,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := convert_hours_2_seconds(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
