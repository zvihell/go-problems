package convertanimals

import "testing"

func TestConvertAnimals(t *testing.T) {
	testCases := []struct {
		name     string
		input1   int
		input2   int
		input3   int
		expected int
	}{
		{
			name:     "test1",
			input1:   2,
			input2:   3,
			input3:   5,
			expected: 36,
		},
		{
			name:     "test2",
			input1:   1,
			input2:   2,
			input3:   3,
			expected: 22,
		},
		{
			name:     "test3",
			input1:   5,
			input2:   2,
			input3:   8,
			expected: 50,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := convert_animals_2_legs(tc.input1, tc.input2, tc.input3)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
