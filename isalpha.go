package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, r := range s {
		if !(('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9')) {
			return false
		}
	}
	return true
}
