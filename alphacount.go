package piscine

func AlphaCount(s string) int {
	compt := 0
	list := []byte(s)
	for i := range list {
		if (list[i] >= 65 && list[i] <= 90) || (list[i] >= 97 && list[i] <= 122) {
			compt++
		}
	}
	return compt
}
