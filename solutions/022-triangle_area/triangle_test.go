package trianglearea

import "testing"

func TestTriangle(t *testing.T) {
	testCases := []struct {
		name     string
		base     int
		height   int
		expected float64
	}{
		{
			name:     "test 1",
			base:     3,
			height:   2,
			expected: 3,
		},
		{
			name:     "test 2",
			base:     7,
			height:   4,
			expected: 14,
		},
		{
			name:     "test 3",
			base:     10,
			height:   10,
			expected: 50,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := triangle_area(tc.base, tc.height)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
