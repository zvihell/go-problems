package flipbool

func flip_bool(b bool) int {
	if b == true {
		return 0
	} else if b == false {
		return 1
	}
	return -1
}
