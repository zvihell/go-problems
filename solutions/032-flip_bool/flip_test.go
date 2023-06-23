package flipbool

import "testing"

func TestFlip(t *testing.T) {
	testCases := []struct {
		name     string
		input    bool
		expected int
	}{
		{
			name:     "test 1",
			input:    true,
			expected: 0,
		},
		{
			name:     "test 1",
			input:    false,
			expected: 1,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := flip_bool(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
