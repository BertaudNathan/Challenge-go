package piscine

import (
	"github.com/01-edu/z01"
)

func AtoiBase(s string, base string) int {
	nb2 := 0

	if Validbase(base) == false {
		return 0
	}
	if s == string(base[0]) {
		z01.PrintRune(rune('0'))
	}
	StrRev(base)
	for i := range s {
		nb2 += RecursivePower(int(s[i]), i)
	}
	return nb2
}
