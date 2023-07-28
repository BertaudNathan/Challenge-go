package piscine

func Index(s string, toFind string) int {
	if StrLen(s) < StrLen(toFind) {
		return -1
	}

	if toFind == "" {
		return 0
	}
	if StrLen(s) == 0 {
		return -1
	}
	for i := range s {
		if StrLen(s[i:]) < StrLen(toFind) {
			return -1
		}
		if s[i] == toFind[0] {
			if s[i:i+StrLen(toFind)] == toFind {
				return i
			}
		}
	}
	return -1
}

/*for j := 0; j <= StrLen(toFind); j++ {
if toFind[j] != s[index] {
	fmt.Println("faux /n")
	break
} else {
	fmt.Println("vrai /n")
	index++
}
return i*/
