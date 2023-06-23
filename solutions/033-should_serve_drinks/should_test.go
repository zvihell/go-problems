package shouldservedrinks

import "testing"

func TestShould(t *testing.T) {
	testCases := []struct {
		name     string
		age      int
		on_break bool
		expected bool
	}{
		{
			name:     "test1",
			age:      17,
			on_break: true,
			expected: false,
		},
		{
			name:     "test2",
			age:      19,
			on_break: false,
			expected: true,
		},
		{
			name:     "test3",
			age:      30,
			on_break: true,
			expected: false,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := should_serve_drinks(tc.age, tc.on_break)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
