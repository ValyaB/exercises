package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func plus(x, y int) int {
	return x + y
}
func mult(x, y int) int {
	return x * y
}
func main() {
	var ints [2][]byte
	var ln [2]int
	var err error
	reader := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {

		if err := reader.Err(); err != nil {
			log.Fatal(err)
		}
		log.Printf("Enter integer N%d:", i+1)
		reader.Scan()
		text := reader.Text()
		_, err = strconv.Atoi(text)
		if err != nil {
			log.Println("please enter digits")
			i--
			continue
		}
		ln[i] = len(text)
		for j := ln[i] - 1; j >= 0; j-- {
			ints[i] = append(ints[i], text[int(j)])
		}
	}

	//log.Println(math.Max(float64(ln[0]), float64(ln[1])))
}
