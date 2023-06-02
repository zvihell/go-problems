package newword

func newWord(s string) string {
	if len(s) == 0 {
		return "Строка пустая"
	}
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return ""
}
