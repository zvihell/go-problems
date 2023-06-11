package footballpoints

import (
	"testing"
)

func TestFootbal(t *testing.T) {
	testCases := []struct {
		name     string
		wins     int
		draws    int
		losses   int
		expected int
	}{
		{
			name:     "test 1",
			wins:     3,
			draws:    4,
			losses:   2,
			expected: 13,
		},
		{
			name:     "test 2",
			wins:     5,
			draws:    0,
			losses:   2,
			expected: 15,
		},
		{
			name:     "test 3",
			wins:     0,
			draws:    0,
			losses:   1,
			expected: 0,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := footbal_points(tc.wins, tc.draws, tc.losses)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
