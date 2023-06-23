package distancehome

func distance_home(nums []int) int {
	total := 0

	for _, items := range nums {
		total += items
	}
	if total < 0 {
		return -total
	}
	return total
}
