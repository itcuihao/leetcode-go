package permutation

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {

	words := [][]int{
		{1, 3, 2},
		// {1, 2, 3},
		// {3, 2, 1},
		// {1, 1, 5},
	}

	for _, word := range words {
		fmt.Println("old:", word)
		nextPermutation(word)
		fmt.Println("new:", word)
	}

}
func TestReverse(t *testing.T) {

	words := []int{1, 2, 3}

	reverse(words, 0, 1)

	fmt.Println(words)
}
