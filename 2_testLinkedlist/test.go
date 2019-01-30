package main

import (
	linkedlist "exercises/2_0linkedlist"
)

func main() {

	f := linkedlist.New(&linkedlist.Node{})
	f2 := linkedlist.New(&linkedlist.Node{})
	// for i := 0; i < 1; i++ {

	// 	f.AddToEnd(&linkedlist.Node{Value: "test" + strconv.Itoa(i)})
	// }
	// for i := 0; i < 2; i++ {

	// 	f.AddToEnd(&linkedlist.Node{Value: "test" + strconv.Itoa(i)})
	// }

	// f.PrintFeed()
	// f.DeDup()

	// f.PrintFeed()
	// fmt.Println(linkedlist.KtoLast(f, 3))
	// f.PrintFeed()
	// f.CutNode("test1")
	// f.PrintFeed()
	// f.CutNode("test5")
	// f.PrintFeed()
	// f.CutNode("test0")
	// f.PrintFeed()
	// fmt.Println("\ntest Partition\n")

	// for i := 3; i < 6; i++ {

	// 	f.AddToEnd(&linkedlist.Node{Value: strconv.Itoa(i)})
	// }
	// for i := 0; i < 6; i++ {

	// 	f.AddToEnd(&linkedlist.Node{Value: strconv.Itoa(i)})
	// }
	// f.PrintFeed()
	// f = f.Partition(5)
	// f.PrintFeed()

	//test Summ of feeds

	// f.AddToEnd(&linkedlist.Node{Value: "9"})
	// f.AddToEnd(&linkedlist.Node{Value: "9"})
	// f.AddToEnd(&linkedlist.Node{Value: "9"})

	// f2.AddToEnd(&linkedlist.Node{Value: "1"})

	// linkedlist.SumFeedsInts(f, f2).PrintFeed()

	// f = linkedlist.New(&linkedlist.Node{})
	// f2 = linkedlist.New(&linkedlist.Node{})

	// f.AddToEnd(&linkedlist.Node{Value: "7"})
	// f.AddToEnd(&linkedlist.Node{Value: "1"})
	// f.AddToEnd(&linkedlist.Node{Value: "6"})
	// f2.AddToEnd(&linkedlist.Node{Value: "5"})
	// f2.AddToEnd(&linkedlist.Node{Value: "9"})
	// f2.AddToEnd(&linkedlist.Node{Value: "2"})

	// linkedlist.SumFeedsInts(f, f2).PrintFeed()

	// f = linkedlist.New(&linkedlist.Node{})
	// f2 = linkedlist.New(&linkedlist.Node{})

	// f.AddToEnd(&linkedlist.Node{Value: "9"})
	// f.AddToEnd(&linkedlist.Node{Value: "9"})
	// f.AddToEnd(&linkedlist.Node{Value: "6"})
	// f2.AddToEnd(&linkedlist.Node{Value: "1"})
	// f2.AddToEnd(&linkedlist.Node{Value: "9"})

	// linkedlist.SumFeedsInts(f, f2).PrintFeed()

	f = linkedlist.New(&linkedlist.Node{})
	f2 = linkedlist.New(&linkedlist.Node{})

	f.AddToEnd(&linkedlist.Node{Value: "1"})
	f.AddToEnd(&linkedlist.Node{Value: "2"})
	f.AddToEnd(&linkedlist.Node{Value: "6"})
	//f2.AddToEnd(&linkedlist.Node{Value: "1"})
	//f2.AddToEnd(&linkedlist.Node{Value: "9"})

	linkedlist.SumFeedsInts(f, f2).PrintFeed()
}
