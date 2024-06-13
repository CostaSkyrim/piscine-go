package piscine

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)
	reversedMenu := make([]string, length)
	for i := range menu {
		reversedMenu[length-1-i] = menu[i]
	}
	return reversedMenu
}
