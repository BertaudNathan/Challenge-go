package piscine

func Index2(s string, toFind string) []int {
	list := []int{0}
	if StrLen(s) < StrLen(toFind) {
		return list
	}

	if toFind == "" {
		return list
	}
	if StrLen(s) == 0 {
		return list
	}
	for i := range s {
		if StrLen(s[i:]) < StrLen(toFind) {
			return list
		}
		if s[i] == toFind[0] {
			if s[i:i+StrLen(toFind)] == toFind {
				list = append(list, i)
			}
		}
	}
	return list
}

func Split(s, sep string) []string {
	listsep := Index2(s, sep)
	liststring := []string{}
	listsep = append(listsep, len(s))
	liststring = append(liststring, s[listsep[0]:listsep[1]])
	for i := 1; i <= len(listsep)-2; i++ {
		liststring = append(liststring, s[listsep[i]+len(sep):listsep[i+1]])
	}
	return liststring
}
