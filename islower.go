package piscine

func IsLower(s string) bool {
	list := []byte(s)
	for i := range list {
		if (list[i] >= 97 && list[i] <= 122) == false {
			return false
		}
	}
	return true
}
