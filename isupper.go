package piscine

func IsUpper(s string) bool {
	list := []byte(s)
	for i := range list {
		if (list[i] >= 65 && list[i] <= 90) == false {
			return false
		}
	}
	return true
}
