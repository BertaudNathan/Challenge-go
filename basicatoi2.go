package piscine

func BasicAtoi2(s string) int {
	tab1 := []rune{}
	nombre := 0
	for i := range s {
		tab1 = append(tab1, rune(s[i]-48))
	}
	for i := range tab1 {
		if tab1[i] != 0 && tab1[i] != 1 && tab1[i] != 2 && tab1[i] != 3 && tab1[i] != 4 && tab1[i] != 5 && tab1[i] != 6 && tab1[i] != 7 && tab1[i] != 8 && tab1[i] != 9 {
			return 0
		}
		nombre = nombre*10 + int(tab1[i])
	}
	return nombre
}
