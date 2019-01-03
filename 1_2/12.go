package main

import (
	"fmt"
	"os"
)

func main() {
	var str1, str2 = "123456789000", "098765400321"
	mapChar := make(map[rune]int)
	if len(str1) != len(str2) {
		fmt.Print("lenght is not equal")
		os.Exit(0)
	}
	for _, val := range str2 {
		if _, ok := mapChar[val]; ok {
			mapChar[val] = mapChar[val] + 1
		} else {
			mapChar[val] = 1
		}
	}

	for _, val := range str1 {
		if valCount, ok := mapChar[val]; ok && (valCount > 0) {
			valCount--
		} else {
			fmt.Printf("there is no permutation %v of %v\n", str1, str2)
			os.Exit(0)
		}

	}
	fmt.Printf("there is permutation %v of %v\n", str1, str2)
}
