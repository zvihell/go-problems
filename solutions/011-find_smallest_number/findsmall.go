package findsmallestnumber

func find_smallest_number(s []int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, item := range s {
		if item < min {
			min = item
		}
	}
	return min
}
