package strstr

func strStr(haystack string, needle string) int {
	lh, ln := len(haystack), len(needle)
	for i := 0; i <= lh; i++ {
		if haystack[i:i+ln] == needle {
			return i
		}
	}
	return -1
}
