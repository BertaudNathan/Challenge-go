package piscine

func Join(strs []string, sep string) string {
	txt := ""
	for i := range strs {
		txt += strs[i]
		txt += sep
	}
	if txt != "" {
		txt = txt[:len(txt)-len(sep)]
	}
	return txt
}
