package piscine

func IterativeFactorial(nb int) int {
	nb2 := 1
	if nb < 0 || nb > 20 {
		return 0
	}
	for i := nb; i >= 1; i-- {
		nb2 = nb2 * i
	}
	return nb2
}
