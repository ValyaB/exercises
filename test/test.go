package solution

import (
	"fmt"
	"strings"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {

	max, count := 0, 0
	position := 0
	if len(S) > 100 {
		S = S[:100]
	}
	fmt.Println(len(S))
	for i, v := range S {
		if v == '!' || v == '.' || v == '?' {
			if i != position {
				count = len(strings.Fields(S[position:i]))
			} else {
				count = 1
			}
			if count > max {
				max = count
			}
			fmt.Println(S[position:i], max, position, i)
			position = i + 1

		}
	}

	return max
}
