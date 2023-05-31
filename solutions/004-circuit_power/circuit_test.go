package circuitpower

import "testing"

func TestCircuitpower(t *testing.T) {
	testCases := []struct {
		name     string
		input1   int
		input2   int
		expected int
	}{
		{
			name:     "test1",
			input1:   230,
			input2:   10,
			expected: 2300,
		},
		{
			name:     "test2",
			input1:   -3,
			input2:   2,
			expected: -6,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := circuit_power(tc.input1, tc.input2)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}

}
