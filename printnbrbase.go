package piscine

import (
	"github.com/01-edu/z01"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Validbase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := range base {
		for j := range base {
			if (base[i] == base[j] && i != j) || base[i] == '-' || base[i] == '+' {
				return false
			}
		}
	}
	return true
}

func PrintNbrBase(nbr int, base string) {
	if Validbase(base) == false {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}
	list := []int{}
	for nbr != 0 {
		rest := nbr % len(base)
		list = append(list, rest)
		nbr = nbr / len(base)
		if nbr == 1 {
			list = append(list, 1)
			break
		}
	}
	for i := len(list) - 1; i >= 0; i-- {
		z01.PrintRune(rune(base[abs(list[i])]))
	}
}
