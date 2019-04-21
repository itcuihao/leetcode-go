package subsequence

func isSubsequence(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls == 0 && lt > 0 {
		return true
	}
	if ls > 0 && lt == 0 {
		return false
	}
	i, j := 0, 0
	for i < ls && j < lt {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == ls
}
