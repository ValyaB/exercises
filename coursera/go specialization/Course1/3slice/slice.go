package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"syscall"
)

func main() {
	var doneChannel = make(chan bool)

	signal.Ignore(syscall.SIGINT)

	go func() {

		slice := make([]int, 3)
		slicetemp := make([]int, 3)
		fmt.Println("Start loop. Type 'X' for exit. \nEnter integer:")
		scanner := bufio.NewScanner(os.Stdin)
		for ii := 0; scanner.Scan(); ii++ {
			fmt.Println("Type 'X' for exit. Enter integer:")
			text := scanner.Text()
			i, err := strconv.Atoi(text)
			if err != nil {
				if text == "X" {
					doneChannel <- true
				}
				ii--
				continue
			}
			if ii < 3 {

				slicetemp[ii] = i

				for i, v := range slicetemp {
					slice[i] = v
				}

			} else {
				slice = append(slice, i)
			}
			sort.Ints(slice)
			fmt.Println(slice)
		}

	}()

	<-doneChannel

}
