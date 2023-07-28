package piscine

func Map(f func(int) bool, a []int) []bool {
	list := []bool{}
	for _, i := range a {
		list = append(list, f(i))
	}
	return list
}
