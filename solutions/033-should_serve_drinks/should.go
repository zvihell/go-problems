package shouldservedrinks

func should_serve_drinks(age int, on_break bool) bool {
	if age < 18 || on_break == true {
		return false
	}
	return true
}
