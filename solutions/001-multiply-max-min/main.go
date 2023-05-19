package multiplymaxmin

func multiplyMaxMin(array []int) int {
	if len(array) == 0 {
		return 0
	}
	min := array[0]
	max := array[0]
	for _, item := range array {
		if min > item {
			min = item
		}
		if max < item {
			max = item
		}
	}
	return min * max
}
