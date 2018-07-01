package prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	p := strs[0]

	for k, v := range p {
		for x, y := range strs {
			if y[k] != byte(v) {
				return strs[x][:k]
			}
		}
	}
	return ""
}
