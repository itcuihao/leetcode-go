package conversion

import (
	"testing"
)

func TestConvert(t *testing.T) {
	s := "abcde"
	t.Log(convert(s, 3))
}
