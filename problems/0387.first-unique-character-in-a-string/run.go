package string

func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	check := make(map[rune]int)
	for _, v := range s {
		check[v]++
	}
	for k, v := range s {
		if check[v] == 1 {
			return k
		}
	}
	return -1
}
