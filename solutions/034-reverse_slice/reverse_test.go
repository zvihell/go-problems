package reverse

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Test 1",
			input:    []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "Test 2",
			input:    []int{10, 20, 30, 40, 50},
			expected: []int{50, 40, 30, 20, 10},
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := revers(tc.input)

		if reflect.DeepEqual(result, tc.expected) {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}

	}

}
