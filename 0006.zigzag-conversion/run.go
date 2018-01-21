package conversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	l := numRows*2 - 2
	ls := len(s)
	ss := ""

	for i := 0; i < numRows; i++ {
		j := i
		for j < ls {
			ss += string(s[j])
			j += l
			if i != 0 && i != numRows-1 && i-2*j < ls {
				ss += string(s[j-2*i])
			}
		}
	}
	return ss
}
