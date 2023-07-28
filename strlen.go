package piscine

func StrLen(txt string) (taille int) {
	len := 0
	for taille := range txt {
		len++
		taille++
	}
	taille = len
	return
}
