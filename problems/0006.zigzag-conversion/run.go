package conversion

import "bytes"

func convert(s string, numRows int) string {
	ls := len(s)
	if numRows == 1 || ls < numRows {
		return s
	}

	buf := bytes.Buffer{}

	l := numRows*2 - 2

	for i := 0; i < ls; i += l {
		buf.WriteByte(s[i])
	}

	for i := 1; i <= numRows-2; i++ {

		buf.WriteByte(s[i])

		for k := l; k-i < len(s); k += l {
			buf.WriteByte(s[k-i])
			if k+i < ls {
				buf.WriteByte(s[k+i])
			}
		}
	}

	for i := numRows - 1; i < ls; i += l {
		buf.WriteByte(s[i])
	}

	return buf.String()
}
