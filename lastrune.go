package piscine

func LastRune(s string) rune {
	list := []rune(s)
	return list[StrLen(s)-1]
}
