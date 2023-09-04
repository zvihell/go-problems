package reverse

func revers(arr []int) []int {
	var reverse []int

	for i := len(arr) - 1; i >= 0; i-- {
		reverse = append(reverse, arr[i])
	}
	return reverse
}
