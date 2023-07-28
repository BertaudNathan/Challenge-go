package piscine

func StrRev(txt string) (rev string) {
	liste := []string{}
	for lettre := range txt {
		liste = append(liste, string(txt[lettre]))
	}
	for i := len(liste) - 1; i >= 0; i-- {
		rev = rev + liste[i]
	}
	return
}
