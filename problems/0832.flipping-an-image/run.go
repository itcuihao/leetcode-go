package image

import "fmt"

func Run() {
	A := [][]int{
		[]int{1, 1, 0},
		[]int{1, 0, 1},
		[]int{0, 0, 0},
	}
	r := flipAndInvertImage(A)
	fmt.Println(r)
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, a := range A {
		for i := 0; i < len(a)/2; i++ {
			a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
		}
		for i := 0; i < len(a); i++ {
			a[i] = a[i] ^ 1
		}
	}
	return A
}
