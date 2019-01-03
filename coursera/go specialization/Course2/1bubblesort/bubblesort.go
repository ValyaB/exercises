package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Swap two elements of slice
func Swap(slice []int, j int) {
	slice[j], slice[j+1] = slice[j+1], slice[j]
}

// BubbleSort of slice
func BubbleSort(intSlice []int) {
	for i := range intSlice {
		for j := 0; j < len(intSlice)-i-1; j++ {
			if intSlice[j] > intSlice[j+1] {
				Swap(intSlice, j)
			}

		}
	}
}

func main() {
	var sliceInt []int
	fmt.Println("Enter sequence of up to 10 integers with space as delimeter:")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	text := reader.Text()
	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range strings.Split(text, " ") {
		ints, err := strconv.Atoi(v)
		if err == nil && len(sliceInt) < 10 {
			sliceInt = append(sliceInt, ints)
		}
	}
	BubbleSort(sliceInt)
	fmt.Println("Sorted sequence:")
	fmt.Printf("%v\n", sliceInt)
}
