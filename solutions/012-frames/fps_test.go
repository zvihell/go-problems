package fps

import "testing"

func TestFps(t *testing.T) {
	testCases := []struct {
		name     string
		input1   int
		input2   int
		expected int
	}{
		{
			name:     "test1",
			input1:   1,
			input2:   1,
			expected: 60,
		},
		{
			name:     "test2",
			input1:   10,
			input2:   1,
			expected: 600,
		},
		{
			name:     "test3",
			input1:   10,
			input2:   25,
			expected: 15000,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := frames(tc.input1, tc.input2)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
