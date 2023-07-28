package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	list := []int{}
	a := ConvertNbsub(n, list)
	SortIntegerTable(a)
	if len(a) == 0 {
		a = append(a, 0)
	}
	for i := range a {
		z01.PrintRune(rune(a[i] + 48))
	}
}

func ConvertNbsub(n int, list []int) (list2 []int) {
	if n == 0 {
	}
	list2 = list
	if n <= 0 {
		return list2
	}
	for i := 0; i <= 9; i++ {
		if (n-i)%10 == 0 {
			list = append(list2, i)
			n = (n - i) / 10
			break
		}
	}
	z01.PrintRune('\n')
	return ConvertNbsub(n, list)
}
