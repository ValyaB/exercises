package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	min := A[0]
	count := 0
	for _, v := range A {
		if min < v {
			min = v
			count++

		} else {
			count--
		}
	}
	if count < 0 {
		count = 0
	}
	return count
}
