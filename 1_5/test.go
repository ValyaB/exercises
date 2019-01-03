package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

var flag, fequal = false, false

func compareString(s1, s2 string) bool {

	if utf8.RuneCountInString(s1) > utf8.RuneCountInString(s2) {
		s1, s2 = s2, s1
		fmt.Println("swap")
	}
	if utf8.RuneCountInString(s1) == utf8.RuneCountInString(s2) {
		fequal = true
		fmt.Println("equal")
	}

	j := 0
	for i := range s2 {
		//fmt.Println(j, utf8.RuneCountInString(s1), s1)
		//fmt.Println(i, utf8.RuneCountInString(s2), s2)
		if j >= utf8.RuneCountInString(s1) {
			return true
		}
		if s2[i] != s1[j] {
			if flag {
				return false
			}
			flag = true
		}
		if s2[i] == s1[j] || fequal {
			_, w := utf8.DecodeRuneInString(s1[j:])
			j = j + w
		}
	}
	return true
}
func main() {
	var str1, str2 = "日本語", "日p本語"

	dif := math.Abs(float64(utf8.RuneCountInString(str1) - utf8.RuneCountInString(str2)))
	if dif == 0 || dif == 1 {
		fmt.Println(compareString(str1, str2))
	} else {
		fmt.Println("false")
	}
}
