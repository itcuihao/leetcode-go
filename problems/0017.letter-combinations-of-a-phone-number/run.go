package closest


var jp = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
func letterCombinations(digits string) []string {
	ld := len(digits)
	// 初始化
	r := []string{}
	if ld == 0 {
		return r
	}
r=[]string{""}
	for i := 0; i < ld; i++ {
		var tmp []string
		// 获取jp对应的下标
		index := digits[i] - '0'
		for _, k := range r {
			for _, j := range jp[index] {
				tmp = append(tmp, k+string(j))
			}
		}
		// 重新给r赋值
		r = tmp
	}
	return r
}
