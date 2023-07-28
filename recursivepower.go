package piscine

func RecursivePower(nb int, power int) int {
	res := 1
	if power < 0 {
		return 0
	}
	if power == 0 {
		return res
	} else {
		return nb * IterativePower(nb, power-1)
	}
}
