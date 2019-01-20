package solution

import (
	"fmt"
	"strconv"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int) int {
	s := strconv.Itoa(A)
	resultS := ""
	lenght := len(s)
	if lenght == 0 || lenght == 1 {
		return A
	}
	for i, j := 0, lenght; i < lenght/2; i++ {

		resultS = resultS + string(s[i]) + string(s[j-1])
		fmt.Println(resultS)
		j--
	}
	if lenght%2 != 0 {
		resultS = resultS + string(s[lenght/2])
	}

	resultI, _ := strconv.Atoi(resultS)
	return resultI
}
