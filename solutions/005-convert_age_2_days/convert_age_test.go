package convertage

import "testing"

func TestConvertAge(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "test1",
			input:    25,
			expected: 9125,
		},
		{
			name:     "test2",
			input:    8,
			expected: 2920,
		},
		{
			name:     "test3",
			input:    0,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := convert_age_2_days(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}

}
