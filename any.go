package piscine

func Any(f func(string) bool, a []string) bool {
	for _, i := range a {
		if f(i) == true {
			return true
		}
	}
	return false
}
