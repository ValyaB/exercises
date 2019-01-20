package main

import (
	"fmt"
	"os"
)

func main() {
	var str1, str2 = "098765400321", "098765400321"
	mapChar := make(map[rune]int)
	if len(str1) != len(str2) {
		fmt.Print("lenght is not equal")
		os.Exit(0)
	}
	for _, val := range str2 {
		//	if _, ok := mapChar[val]; ok {
		mapChar[val]++
		//
	}

	for _, val := range str1 {
		if _, ok := mapChar[val]; !ok {
			fmt.Printf("there is no permutation %v of %v\n", str1, str2)
			os.Exit(0)
		}
		mapChar[val]--
		if mapChar[val] < 0 {
			fmt.Printf("there is no permutation %v of %v\n", str1, str2)
			os.Exit(0)
		}

	}
	fmt.Printf("there is permutation %v of %v\n", str1, str2)

}
