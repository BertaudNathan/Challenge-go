package piscine

func ToLower(s string) string {
	list := []byte(s)
	s2 := ""
	for i := range list {
		if list[i] >= 65 && list[i] <= 90 {
			list[i] = list[i] + 32
		}
		s2 += string(rune(list[i]))
	}
	s = s2
	return s
}
