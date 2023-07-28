package piscine

func AlphaCount2(s string) int {
	compt := 0
	list := []byte(s)
	for i := range list {
		if (list[i] >= 65 && list[i] <= 90) || (list[i] >= 97 && list[i] <= 122) || (list[i] >= 48 && list[i] <= 57) {
			compt++
		}
	}
	return compt
}

func IsAlpha(s string) bool {
	if StrLen(s) == AlphaCount2(s) {
		return true
	}
	return false
}
