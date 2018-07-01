package integer

// 到着遍历？
func romanToInt(s string) int {
	k := 0
	ls := len(s)
	if ls == 0 {
		return k
	}
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	getRoman := func(b byte) int {
		if v, ok := roman[b]; ok {
			return v
		}
		return 0
	}

	tn := 0
	for i := ls - 1; i >= 0; i-- {
		tmp := getRoman(s[i])

		if tmp < tn {
			k -= tmp
		} else {
			k += tmp
		}

		tn = tmp

	}
	return k
}
