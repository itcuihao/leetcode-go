package prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	p := strs[0]

	for k, v := range p {
		for j := 0; j < len(strs); j++ {
			if strs[j][k] != byte(v) {
				return strs[j][:k]
			}
		}
	}
	return p
}
