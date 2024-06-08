package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	totalLen := 0
	for i, arg := range args {
		totalLen += len(arg)
		if i < len(args)-1 {
			totalLen++
		}
	}

	result := make([]byte, totalLen)
	pos := 0

	for i, arg := range args {
		for j := 0; j < len(arg); j++ {
			result[pos] = arg[j]
			pos++
		}
		if i < len(args)-1 {
			result[pos] = '\n'
			pos++
		}
	}
	return string(result)
}
