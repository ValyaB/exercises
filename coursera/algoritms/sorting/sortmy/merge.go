package sortmy

import (
	"sync"
)

func merge(l, r []int) []int {

	var newslice []int
	//fmt.Println("merge:", newslice, l, r, len(r), len(l))
	//time.Sleep(2 * time.Second)
	for len(r) > 0 || len(l) > 0 {
		switch {
		case len(r) == 0:
			return append(newslice, l...)

		case len(l) == 0:
			return append(newslice, r...)

		case l[0] <= r[0]:
			newslice = append(newslice, l[0])
			l = l[1:]

		case l[0] > r[0]:
			newslice = append(newslice, r[0])
			r = r[1:]
		}

	}
	//fmt.Println(newslice, l, r)
	return newslice
}

var wg sync.WaitGroup

func PARALLELMergeSort(data []int) []int {
	sem := make(chan struct{}, 4)
	return multiMergeSort(data, sem)
}

func multiMergeSort(data []int, sem chan struct{}) []int {
	l := len(data)
	if l < 2 {
		return data
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	var left, right []int
	select {
	case sem <- struct{}{}:
		go func() {
			left = multiMergeSort(data[:l/2], sem)
			<-sem
			wg.Done()
		}()
	default:
		left = SimpleMergeSort(data[:l/2])
		wg.Done()
	}
	select {
	case sem <- struct{}{}:
		go func() {
			right = multiMergeSort(data[l/2:], sem)
			<-sem
			wg.Done()
		}()
	default:
		right = SimpleMergeSort(data[l/2:])
		wg.Done()
	}
	wg.Wait()
	return merge(left, right)
}

func SimpleMergeSort(data []int) []int {
	l := len(data)
	if l <= 1 {
		return data
	}

	left := SimpleMergeSort(data[:l/2])
	right := SimpleMergeSort(data[l/2:])
	return merge(left, right)
}

func MMergeSort(data []int, res chan []int) {
	if len(data) < 2 {
		res <- data
		return
	}

	leftChan := make(chan []int)
	rightChan := make(chan []int)
	middle := len(data) / 2

	go MMergeSort(data[:middle], leftChan)
	go MMergeSort(data[middle:], rightChan)

	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	res <- merge(ldata, rdata)
	return
}

func RunMMergeSort(data []int) (multiResult []int) {
	res := make(chan []int)
	go MMergeSort(data, res)
	multiResult = <-res
	return
}
