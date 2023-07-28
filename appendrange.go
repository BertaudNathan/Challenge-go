package piscine

func AppendRange(min, max int) []int {
	list := []int{}
	if min == max || min > max {
		return []int(nil)
	}
	for i := min; i < max; i++ {
		list = append(list, i)
	}
	return list
}
