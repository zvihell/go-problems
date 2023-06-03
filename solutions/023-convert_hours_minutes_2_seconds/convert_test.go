package converthoursminutes

import "testing"

func TestConvertHoursMinutes(t *testing.T) {
	testCases := []struct {
		name     string
		hours    int
		minutes  int
		expected int
	}{
		{
			name:     "test1",
			hours:    1,
			minutes:  3,
			expected: 3780,
		},
		{
			name:     "test2",
			hours:    2,
			minutes:  0,
			expected: 7200,
		},
		{
			name:     "test3",
			hours:    0,
			minutes:  0,
			expected: 0,
		},
	}
	for _, tc := range testCases {
		println("Проверка теста " + tc.name + "..")
		result := convert_hours_minutes_2_seconds(tc.hours, tc.minutes)
		if result == tc.expected {
			println("✅ Тест «" + tc.name + "» пройден")
		} else {
			t.Errorf("🤬 Тест «%v» провален, получено %v, ожидалось %v\n", tc.name, result, tc.expected)
		}
	}
}
