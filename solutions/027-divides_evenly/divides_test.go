package dividesevenly

import "testing"

func TestDivides(t *testing.T) {
	testCases := []struct {
		name     string
		max      int
		min      int
		expected bool
	}{
		{
			name:     "test 1",
			max:      98,
			min:      7,
			expected: true,
		},
		{
			name:     "test 2",
			max:      85,
			min:      4,
			expected: false,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := divides_evenly(tc.max, tc.min)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
