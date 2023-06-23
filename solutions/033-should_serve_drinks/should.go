package shouldservedrinks

func should_serve_drinks(age int, on_break bool) bool {
	if age < 18 {
		return false
	}
	if on_break {
		return false
	}
	return true
}
