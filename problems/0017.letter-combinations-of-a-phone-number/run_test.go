package closest

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {

	target := "23"

	s := letterCombinations(target)
	fmt.Println(s)
}
