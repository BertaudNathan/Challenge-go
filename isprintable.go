package piscine

func IsPrintable(s string) bool {
	for i := 0; i <= StrLen(s)-1; i++ {
		if s[i] <= 31 || s[i] == 127 {
			return false
		}
	}
	return true
}
