package piscine

func CountIf(f func(string) bool, tab []string) int {
	compt := 0
	for _, i := range tab {
		if f(i) == true {
			compt++
		}
	}
	return compt
}
