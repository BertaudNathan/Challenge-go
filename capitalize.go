package piscine

func Capitalize(s string) string {
	list := []rune(s)
	s2 := ""
	if list[0] >= 97 && list[0] <= 122 {
		list[0] = list[0] - 32
	}
	for i := 1; i <= len(list)-1; i++ {
		if i == len(list)-1 && ((list[len(list)-1] >= 65 && list[len(list)-1] <= 90) || (list[len(list)-1] >= 97 && list[len(list)-1] <= 122) || (list[len(list)-1] >= 48 && list[len(list)-1] <= 57)) == false {
			break
		}
		if ((list[i] >= 65 && list[i] <= 90) || (list[i] >= 97 && list[i] <= 122) || (list[i] >= 48 && list[i] <= 57)) == false && (list[i+1] >= 97 && list[i+1] <= 122) {
			list[i+1] = list[i+1] - 32
		} else if list[i] >= 65 && list[i] <= 90 && ((IsAlpha(string(list[i-1]))) == true || (rune(list[i-1]) >= 48 && rune(list[i-1]) <= 57)) {
			list[i] = list[i] + 32
		}
	}
	for i := range list {
		s2 += string(list[i])
	}
	return s2
}
