package piscine

func MakeRange(min, max int) []int {
	if min == max || min > max {
		return []int(nil)
	}
	list := make([]int, max-min)
	index := 0
	for i := min; i < max; i++ {
		list[index] = i
		index++
	}
	return list
}
