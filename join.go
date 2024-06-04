package piscine

func Join(strs []string, sep string) string {
	result := ""
	for _, str := range strs {
		result = result + str + sep
	}
	return result
}
