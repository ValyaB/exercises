package linkedlist

import (
	"fmt"
)

type Feed struct {
	size   int
	firstN *Node
	lastN  *Node
}
type Node struct {
	Value string
	next  *Node
	prev  *Node
}

func New(first *Node) *Feed {
	var f = Feed{size: 0}
	if first.Value != "" {
		f.size = 1
		f.firstN = first
		f.lastN = first
		fmt.Println("first is not empty")
	}
	return &f
}

func (f *Feed) AddToEnd(n *Node) {
	if f.size == 0 {
		f.firstN, f.lastN = n, n
	}
	if f.size != 0 {
		n.prev = f.lastN
		f.lastN.next = n
		f.lastN = n
	}
	f.size++
}

func (f *Feed) PrintFeed() {
	for n := f.firstN; n != nil; n = n.next {
		fmt.Println(n)
	}
	fmt.Println("size: ", f.size)
}

func (f *Feed) DeDup() {
	mapval := make(map[string]int)
	for n := f.firstN; n != nil; n = n.next {
		if val, ok := mapval[n.Value]; ok && val == 1 {
			//fmt.Println(n.prev.next, n.next)
			if n.next == nil {
				n.prev.next = nil
				f.size--
			} else {
				n.prev.next = n.next
				n.next.prev = n.prev
				f.size--
			}
		} else {
			mapval[n.Value]++
		}
	}

}
