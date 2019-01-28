package main

import (
	linkedlist "exercises/2_0linkedlist"
	"fmt"
	"strconv"
)

func main() {

	f := linkedlist.New(&linkedlist.Node{})

	for i := 0; i < 1; i++ {

		f.AddToEnd(&linkedlist.Node{Value: "test" + strconv.Itoa(i)})
	}
	for i := 0; i < 6; i++ {

		f.AddToEnd(&linkedlist.Node{Value: "test" + strconv.Itoa(i)})
	}

	f.PrintFeed()
	f.DeDup()

	f.PrintFeed()
	fmt.Println(linkedlist.KtoLast(f, 3))
	f.PrintFeed()
	f.CutNode("test1")
	f.PrintFeed()
	f.CutNode("test5")
	f.PrintFeed()
	f.CutNode("test0")
	f.PrintFeed()
}
