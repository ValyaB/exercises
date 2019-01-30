package linkedlist

import (
	"fmt"
	"strconv"
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

func KtoLast(f *Feed, k int) []string {
	//k=1 last element
	arrNode := []string{}
	for n := f.lastN; n.prev != nil; n = n.prev {
		if k > 0 {
			arrNode = append(arrNode, n.Value)
			k--
		}
	}
	return arrNode
}

func (f *Feed) CutNode(cutN string) {

	for n := f.firstN; n != nil; n = n.next {
		if cutN == n.Value {
			if f.size == 1 {
				f.size, f.firstN, f.lastN = 0, nil, nil
				return
			}

			if n.next != nil {
				n.next.prev = n.prev
			} else {
				n.prev.next = nil
				f.lastN = n.prev
			}

			if n.prev != nil {
				n.prev.next = n.next
			} else {
				n.next.prev = nil
				f.firstN = n.next
			}

			f.size--
		}
	}
}

func (f *Feed) Partition(p int) *Feed {
	var head, tail = Feed{}, Feed{}
	for n := f.firstN; n != nil; n = n.next {
		i, err := strconv.Atoi(n.Value)
		if err != nil {
			fmt.Println("ERROR")
		}
		if i < p {
			head.AddToEnd(n)
		} else {
			tail.AddToEnd(n)
		}
	}

	// for n := tail.firstN; n != nil; n = n.next {
	// 	head.AddToEnd(n)
	// }
	if tail.size != 0 {
		tail.firstN.prev = head.firstN
		head.lastN.next = tail.firstN
		head.size = head.size + tail.size
	}
	return &head
}

func SumFeedsInts(f1, f2 *Feed) *Feed {
	var sumF Feed
	var val1, val2, carry int
	for n1, n2 := f1.firstN, f2.firstN; n1 != nil && n2 != nil; {
		val1, _ = strconv.Atoi(n1.Value)
		val2, _ = strconv.Atoi(n2.Value)
		valSum := val1 + val2 + carry

		if valSum < 10 {
			sumF.AddToEnd(&Node{Value: strconv.Itoa(valSum)})
			carry = 0
		} else {
			sumF.AddToEnd(&Node{Value: strconv.Itoa(valSum % 10)})
			carry = valSum / 10
		}

		if f1.size > f2.size && n2.next == nil {
			n1 = n1.next
			n2 = &Node{Value: "0"}
			continue
		}
		if f1.size < f2.size && n1.next == nil {
			n2 = n2.next
			n1 = &Node{Value: "0"}
			continue
		}

		n2 = n2.next
		n1 = n1.next

	}

	if carry != 0 {
		if carry < 10 {
			sumF.AddToEnd(&Node{Value: strconv.Itoa(carry)})
		} else {
			sumF.AddToEnd(&Node{Value: strconv.Itoa(carry % 10)})
			sumF.AddToEnd(&Node{Value: strconv.Itoa(carry / 10)})
		}
	}
	if f1.size == 0 {
		return f2
	}
	if f2.size == 0 {
		return f1
	}
	return &sumF
}
