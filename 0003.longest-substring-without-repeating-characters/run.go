package characters

// The above solution requires at most 2n steps. In fact, it could be optimized to require only n steps. Instead of using a set to tell if a character exists or not, we could define a mapping of the characters to its index. Then we can skip the characters immediately when we found a repeated character.

// The reason is that if s[j]s[j] have a duplicate in the range [i, j)[i,j) with index j'j
// ​′
// ​​ , we don't need to increase ii little by little. We can skip all the elements in the range [i, j'][i,j
// ​′
// ​​ ] and let ii to be j' + 1j
// ​′
// ​​ +1 directly.
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m, max, i := make(map[rune]int, len(s)), 0, 0
	for k, v := range s {
		if j, ok := m[v]; ok {
			i = mmax(j, i)
		}
		max = mmax(max, k-i+1)
		m[v] = k + 1
	}
	return max
}

func mmax(ml, m int) int {
	if ml > m {
		return ml
	}
	return m
}
