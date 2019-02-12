package main

import "fmt"

func RecursiveChannel(number int, product int, result chan int) {

	product = product + number

	if number == 1 {

		result <- product
		return
	}

	go RecursiveChannel(number-1, product, result)
}

func main() {

	result := make(chan int)

	RecursiveChannel(7654321, 0, result)
	answer := <-result

	fmt.Printf("Recursive: %d\n", answer)
}
