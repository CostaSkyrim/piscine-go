package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	started := false
	for _, c := range s {
		if c == '-' && !started {
			sign = -1
			started = true
		} else if c >= '0' && c <= '9' {
			started = true
			result = result*10 + int(c-'0')
		} else if started {
			break
		}
	}
	return sign * result
}

func atoi(s string) int {
	start := 0
	sign := 1
	// Find the index of the first numeric character or '-'
	for i, c := range s {
		if c == '-' || (c >= '0' && c <= '9') {
			if c == '-' {
				sign = -1
			} else {
				start = i
			}
			break
		}
	}
	// Extract the substring containing only numeric characters
	substr := s[start:]
	return sign * TrimAtoi(substr)
}
