package rand7

import (
	"math/rand"
	"time"
)

func rand10() int {
	// num := 41
	// for num >= 40 {
	// 	num = (rand7()-1)*7 + rand7() - 1
	// }
	// return num%10 + 1

	a := rand7()
	b := rand7()

	for {
		if b <= 4 {
			return a
		} else if a <= 4 {
			return b + 3
		}
		a = rand7()
		b = rand7()
	}
}

func rand7() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(7)
}
