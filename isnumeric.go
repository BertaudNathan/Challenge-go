package piscine

func IsNumeric(s string) bool {
	list := []byte(s)
	for i := range list {
		if (list[i] >= 48 && list[i] <= 57) == false {
			return false
		}
	}
	return true
}
