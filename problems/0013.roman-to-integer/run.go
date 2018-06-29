package integer

import "fmt"

// 到着遍历？
func romanToInt(s string) int {
	k := 0
	if len(s) == 0 {
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

	if len(s) == 1 {
		k = getRoman(s)
		return k
	}
	tmap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		tmp := getRoman(string(s[i]))
		fmt.Println("/", tmap)
		fmt.Println("/", tmp)

		if len(tmap) > 0 {

			if v, ok := tmap[string(s[i])]; ok {
				tmp -= v
				fmt.Println(tmp)
				k = tmp
			} else {

				k += tmp
			}

			tmap = make(map[string]int)
		} else {
			k = tmp
		}

		fmt.Println("k:", k)
		switch s[i] {
		case 'I':
			tmap["V"] = 1
			tmap["X"] = 1
		case 'X':
			tmap["L"] = 10
			tmap["C"] = 10
		case 'C':
			tmap["D"] = 100
			tmap["M"] = 100
		}

	}
	return k
}
