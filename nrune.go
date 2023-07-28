package piscine

func NRune(s string, n int) rune {
	if StrLen(s) == 0 || n > StrLen(s) || n < 1 {
		return 0
	}
	list := []rune(s)
	return list[n-1]
}
