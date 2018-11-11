package words

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {

	// s := "wordgoodgoodgoodbestword"
	// words := []string{"word", "good", "best", "good"}
	// s := "barfoothefoobarman"
	// words := []string{"foo", "bar"}
	s := "foobarfoobar"
	words := []string{"foo", "bar"}

	i := findSubstring(s, words)
	fmt.Println(i)
}

func TestArray(t *testing.T) {
	var a array
	words := []string{"a", "b", "c"}
	a.permutation(words, 0, len(words))
	fmt.Println(a.eles)
}
