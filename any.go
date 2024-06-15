package piscine

func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
        	if f(v) {
        		return true
        	}
	}
   	return false
}

func isNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
