package issafebridge

import (
	"testing"
)

func TestSafeBridge(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "test 1",
			input:    "",
			expected: false,
		},

		{
			name:     "test 2",
			input:    "####",
			expected: true,
		},
		{
			name:     "test 3",
			input:    "## ####",
			expected: false,
		},
		{
			name:     "test 4",
			input:    "#",
			expected: true,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := is_safe_bridge(tc.input)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
