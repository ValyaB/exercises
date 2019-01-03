package main

import (
	"fmt"
)

func main() {
	var str1 = "11a7 7"
	runCount := make(map[rune]int)
	odd := 0
	for _, val := range str1 {
		if val != ' ' {
			if _, ok := runCount[val]; ok {
				runCount[val]++
			} else {
				runCount[val] = 1
			}
		}
	}
	for _, val := range runCount {
		if val%2 == 1 {
			odd++
		}
		if odd > 1 {
			fmt.Println("not palindrome", runCount, odd)
			break
		}

	}
	if odd <= 1 {
		fmt.Println("true", runCount, odd)
	}
}
