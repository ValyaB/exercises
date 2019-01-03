package main

import (
	"fmt"
	"os"
)

func main() {
	var str1 = "45qwertyuiop[asdfghjkl;'zxcvbnm,./"
	mapChar := make(map[rune]bool)

	for _, val := range str1 {
		if _, ok := mapChar[val]; ok {
			fmt.Println("not unique")
			os.Exit(0)

		} else {
			mapChar[val] = true
		}
	}
	fmt.Println("unique")
}
