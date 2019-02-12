package main

import "fmt"

func TailRecursive(number int, product int) int {

	product = product + number

	if number == 1 {

		return product
	}

	return TailRecursive(number-1, product)
}

func main() {

	answer := TailRecursive(7654321, 0)
	fmt.Printf("Recursive: %d\n", answer)
}
