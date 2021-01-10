package cookies

import "fmt"

func Run() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	res := findContentChildren(g, s)
	fmt.Println(res)
}

func findContentChildren(g []int, s []int) int {
	children := bubbleSort(g)
	cookies := bubbleSort(s)
	child, cookie := 0, 0
	for child < len(children) && cookie < len(cookies) {
		if children[child] <= cookies[cookie] {
			child++
		}
		cookie++
	}
	return child
}

func bubbleSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
