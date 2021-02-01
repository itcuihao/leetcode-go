package swap

import "fmt"

func Run() {
	A := []int{1, 2, 5}
	B := []int{2, 4}
	res := fairCandySwap(A, B)
	fmt.Println(res)
}

func fairCandySwap(A []int, B []int) []int {
	aSum := sum(A)
	bSum := sum(B)
	diff := aSum - bSum

	res := make([]int, 2)
	checkB := make(map[int]bool, len(B))
	for _, i := range B {
		checkB[i] = true
	}
	for i := 0; i < len(A); i++ {
		res[0] = A[i]
		res[1] = A[i] - diff/2
		if checkB[res[1]] {
			return res
		}
	}
	return nil
}

func sum(arr []int) int {
	s := 0
	for _, i := range arr {
		s += i
	}
	return s
}
