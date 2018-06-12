package zigzag_conversion

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	temp := [][]byte{}
	groupRows := numRows + numRows - 2
	for i := 0; i < len(s); i += groupRows {
		if i+groupRows > len(s) {
			temp = append(temp, convertSubString(s[i:], numRows)...)
		} else {
			temp = append(temp, convertSubString(s[i:i+groupRows], numRows)...)
		}
	}

	sb := bytes.NewBuffer(nil)
	for j := 0; j < numRows; j++ {
		for i := 0; i < len(temp); i++ {
			c := temp[i][j]
			if c > 0 {
				sb.WriteByte(c)
			}
		}
	}

	return sb.String()
}

func convertSubString(sub string, numRows int) [][]byte {
	res := [][]byte{}

	firstLine := make([]byte, numRows)
	copy(firstLine, []byte(sub))
	res = append(res, firstLine)

	if len(sub) > numRows {
		for i := 1; i < numRows-1 && i < len(sub)-numRows+1; i++ {
			tmp := make([]byte, numRows)
			tmp[numRows-i-1] = []byte(sub)[numRows+i-1]
			res = append(res, tmp)
		}
	}

	return res
}
