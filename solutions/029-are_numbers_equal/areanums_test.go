package arenumbersequal

import "testing"

func TestEqual(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected bool
	}{
		{
			name:     "test 1",
			a:        4,
			b:        8,
			expected: false,
		},
		{
			name:     "test 2",
			a:        2,
			b:        2,
			expected: true,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := are_numbers_equal(tc.a, tc.b)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}

}
