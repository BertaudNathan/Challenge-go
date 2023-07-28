package piscine

func Atoi(s string) int {
	state := true
	compteur := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '+' {
			state = true
			compteur++
			break
		}
		if s[i] == '-' {
			state = false
			compteur++
			break
		}
	}
	result := Atoi2(s[compteur:])
	if state == true {
		return result
	}
	if state == false {
		return -1 * result
	}
	return result
}

func Atoi2(s string) int {
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
