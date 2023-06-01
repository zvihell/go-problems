package getfirstelement

func get_first_element(a []int) int {
	if len(a) == 0 {
		return 0
	}
	first := a[0]
	return first
}
