package lessthan100

func less_than_100(a, b int) bool {
	res := a + b
	if res < 100 {
		return true
	}
	return false
}
