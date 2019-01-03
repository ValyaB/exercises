package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type nameFild [20]rune

type person struct {
	fname nameFild
	lname nameFild
}

func main() {

	var p1 []person
	fmt.Println("Enter full path to file:")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	path := reader.Text()

	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		var person1 person

		for i, v := range words {
			if i == 0 {
				for ii, r := range v {
					if ii < 20 {
						//fmt.Printf("%#U starts at byte position %d\n", r, ii)
						person1.fname[ii] = r
					}
				}
			}
			if i == 1 {
				for ii, r := range v {
					if ii < 20 {
						person1.lname[ii] = r
					}
				}
			}
		}
		p1 = append(p1, person1)
	}
	for _, v := range p1 {
		fmt.Println("Fname:", string(v.fname[:]), "Lname:", string(v.lname[:]))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
