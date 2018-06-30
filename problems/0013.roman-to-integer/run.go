package integer

import "fmt"

// 到着遍历？
func romanToInt(s string) int {
	k := 0
	ls := len(s)
	if ls == 0 {
		return k
	}
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	getRoman := func(k string) int {
		if v, ok := roman[k]; ok {
			return v
		}
		return 0
	}

	if ls == 1 {
		k = getRoman(s)
		return k
	}

	tn := 0
	for i := ls - 1; i >= 0; i-- {
		tmp := getRoman(string(s[i]))

		fmt.Println("/", tmp)

		if tmp < tn {
			k -= tmp
		} else {
			k += tmp
		}

		tn = tmp

	}
	return k
}
