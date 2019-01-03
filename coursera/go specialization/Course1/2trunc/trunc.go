package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter floating number or CTRL^C for exit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		ucl := scanner.Text()

		_, err := strconv.ParseFloat(ucl, 64)

		if err == nil {
			i := strings.Split(ucl, ".")
			fmt.Println(i[0])
		} else {
			fmt.Println("It is not float. Please enter floating number (ex: 3.14)")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
