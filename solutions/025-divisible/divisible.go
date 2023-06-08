package divisible

func divisible(number int) bool {
	if number%100 == 0 {
		return true
	} else if number%100 != 100 {
		return false
	}
	return false
}
