package getmiddleelement

func get_middle_element(a []int) int {
	if len(a) == 0 {
		return 0
	}
	index := len(a) / 2
	element := a[index]
	return element
}
