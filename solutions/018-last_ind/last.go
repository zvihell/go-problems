package lastind

func last_ind(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	last := arr[len(arr)-1]
	return last
}
