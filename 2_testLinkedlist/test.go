package main

import (
	linkedlist "exercises/2_0linkedlist"
	"strconv"
)

func main() {

	f := linkedlist.New(&linkedlist.Node{})

	for i := 0; i < 1; i++ {

		f.AddToEnd(&linkedlist.Node{Value: "test" + strconv.Itoa(i)})
	}
	for i := 0; i < 3; i++ {

		f.AddToEnd(&linkedlist.Node{Value: "test" + strconv.Itoa(i)})
	}
	for i := 2; i < 5; i++ {

		f.AddToEnd(&linkedlist.Node{Value: "test" + strconv.Itoa(i)})
	}
	f.PrintFeed()
	f.DeDup()
	f.PrintFeed()

}
