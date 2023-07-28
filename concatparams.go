package piscine

func ConcatParams(args []string) string {
	txt := ""
	for i := range args {
		txt += args[i]
		txt += "\n"
	}
	return txt[0 : len(txt)-1]
}
