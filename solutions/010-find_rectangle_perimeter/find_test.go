package findrectangleperimeter

import "testing"

func TestFindTriangle(t *testing.T) {
	testCases := []struct {
		name     string
		input1   int
		input2   int
		expected int
	}{
		{
			name:     "test1",
			input1:   6,
			input2:   7,
			expected: 26,
		},
		{
			name:     "test2",
			input1:   20,
			input2:   10,
			expected: 60,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := find_rectangle_perimeter(tc.input1, tc.input2)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
