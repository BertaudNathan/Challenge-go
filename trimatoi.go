package piscine

func TrimAtoisub(s string) int {
	list := []string{}
	compt := 0
	compte := 0
	txt := ""
	state := true
	for i := 0; i < len(s); i++ {
		if byte(s[i]) >= 48 && byte(s[i]) <= 57 {
			break
		}
		if s[i] == '+' {
			state = true
			compt++
			break
		}
		if s[i] == '-' {
			state = false
			compt++
			break
		}
	}
	for j := range s {
		if (byte(s[j]) < 48 || byte(s[j]) > 57) == false {
			list = append(list, string(s[j]))
		}
	}
	for i := range list {
		if list[i] != "0" {
			break
		}
		compte++
	}
	list = list[compte:]
	for i := range list {
		txt += list[i]
	}
	if state == true {
		return Atoi(txt)
	} else {
		return -1 * Atoi(txt)
	}
}

func TrimAtoi(s string) int {
	a := TrimAtoisub(s)
	return a
}
