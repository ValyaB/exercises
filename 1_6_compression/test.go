package mycompression

import (
	"bytes"
	"strconv"
)

func mycompression(s string) string {
	var newS bytes.Buffer
	j := 0
	var prevChar rune
	for i, v := range s {
		if v == prevChar || i == 0 {
			j++
		}

		if i != 0 && v != prevChar {
			newS.WriteString(string(prevChar))
			newS.WriteString(strconv.Itoa(j))
			j = 1
		}
		prevChar = v

		if i == len(s)-1 {
			newS.WriteString(string(prevChar))
			newS.WriteString(strconv.Itoa(j))
		}
	}

	return newS.String()
}
