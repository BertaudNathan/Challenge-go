package piscine

func ToUpper(s string) string {
	list := []byte(s)
	s2 := ""
	for i := range list {
		if list[i] >= 97 && list[i] <= 122 {
			list[i] = list[i] - 32
		}
		s2 += string(rune(list[i]))
	}
	s = s2
	return s
}
