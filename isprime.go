package piscine

/*func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	ban := []int{}
	for i := 2; i <= nb-1; i++ {
		if nb%i == 0 {
			ban = append(ban, i)
		}
	}
	if len(ban) == 0 {
		return true
	}
	return false
}
*/

func IsPrime(nb int) bool {
	reste := 0
	flag := true
	if nb < 0 {
		return false
	} else if nb == 0 {
		return false
	} else if nb == 1 {
		return false
	} else {
		for i := 2; i <= nb/2; i++ {
			reste = nb % i
			if reste == 0 {
				flag = false
				break
			}
		}
		return flag
	}
}
