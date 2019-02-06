package main

import (
	"fmt"
	"os"
)

func main() {
	var newS = " 1a2a 1"
	runCount := make(map[rune]int)
	var str1 string
	for _, val := range newS {
		if val != ' ' {
			str1 = str1 + string(val)
		}
	}

	l := len(str1)
	i := 1

	for _, val := range str1 {

		if l/2+l%2 > l/2 && i == l/2+l%2 {
			i++
			continue
		}
		if i <= l/2 {

			if _, ok := runCount[val]; ok {
				runCount[val]++
			} else {
				runCount[val] = 1
			}
		}

		if i > l/2+l%2 {

			if _, ok := runCount[val]; ok {
				runCount[val]--
				if runCount[val] < 0 {
					fmt.Println("not palindrome")
					os.Exit(0)
				}
			} else {
				fmt.Println("not palindrome")
				os.Exit(0)
			}
		}
		i++
	}
	fmt.Println("true")
}
