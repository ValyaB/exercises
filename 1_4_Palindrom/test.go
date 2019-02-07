package main

import (
	"fmt"
)

func main() {

	var str1 = "2    32"

	l := len(str1)
	if l == 0 {
		fmt.Println("empty")
		return
	}

	for t, h := l-1, 0; h < l && t > 0 && t > h; t, h = t-1, h+1 {

		for ; str1[t] == 32; t-- {
			if t <= 0 {
				break
			}
		}
		for ; str1[h] == 32; h++ {
			if h > l {
				break
			}
		}

		if t <= 0 || h > l || t < h {
			break
		}
		if str1[t] != str1[h] {
			fmt.Println("false")
			return
		}
	}

	fmt.Println("true")
}
