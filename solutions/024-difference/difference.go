package difference

func difference(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	max := a[0]
	for _, item := range a {
		if min < item {
			min = item
		}
		if max > item {
			max = item
		}
	}
	return min - max
}
