package bits

import "fmt"

func Run() {
	bits := countBits(5)
	fmt.Println(bits)
}

func countBits(num int) []int {
	bits := make([]int, num+1)
	for i := 0; i <= num; i++ {
		bits[i] = oneCount(i)
	}
	return bits
}

func oneCount(n int) int {
	c := 0
	for ; n > 0; n &= n - 1 {
		c++
	}
	return c
}
