package envelopes

import "sort"

func Run() {
	r := [][]int{
		[]int{5, 4},
		[]int{6, 4},
		[]int{6, 7},
		[]int{2, 3},
	}
	me := maxEnvelopes(r)
	println(me)
}
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	f := []int{}
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}
