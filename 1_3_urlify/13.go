package main

import (
	"fmt"
)

func main() {
	var str1 = "13467   "
	var Out string
	var len = 7

	for i, val := range str1 {
		if val == ' ' && i < len {
		} else {
			Out = Out + string(val)
		}
	}
	fmt.Printf("input: %v\nOutput: %v\n", str1, Out)

}
