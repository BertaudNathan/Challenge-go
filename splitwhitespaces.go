package piscine

func SplitWhiteSpaces(s string) []string {
	list := []int{0}
	list2 := []string{}
	listfinal := []string{}
	for i := range s {
		if string(s[i]) == " " || string(s[i]) == "		" || string(s[i]) == "\n" {
			list = append(list, i)
		}
	}
	list = append(list, len(s))
	list2 = append(list2, s[list[0]:list[1]])
	for i := 1; i <= len(list)-2; i++ {
		list2 = append(list2, s[list[i]+1:list[i+1]])
	}
	for i := range list2 {
		if list2[i] != "" {
			listfinal = append(listfinal, list2[i])
		}
	}
	return listfinal
}
