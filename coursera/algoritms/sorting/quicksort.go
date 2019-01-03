package main

import (
	"bufio"
	"coursera/algoritms/sorting/sortmy"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func sortQuickRand(array []int) {
	//fmt.Println("len", len(array), array)
	if len(array) > 1 {
		rndm := rand.Intn(len(array) - 1)
		array[len(array)-1], array[rndm] = array[rndm], array[len(array)-1]
		pIndex := partition(array)
		//fmt.Println("pIndex", pIndex, array[pIndex], array, array[:pIndex], array[pIndex+1:])

		sortQuickRand(array[:pIndex])
		sortQuickRand(array[pIndex+1:])

	}
}

func sortQuick(array []int) {
	//fmt.Println("len", len(array), array)
	if len(array) > 1 {
		pIndex := partition(array)
		//fmt.Println("pIndex", pIndex, array[pIndex], array, array[:pIndex], array[pIndex+1:])

		sortQuick(array[:pIndex])
		sortQuick(array[pIndex+1:])

	}
}
func partition(a []int) (pIndex int) {

	pivot := a[len(a)-1]
	//fmt.Println("pivot:", pivot)
	pIndex = -1
	for i := range a {
		if a[i] < pivot {
			pIndex++
			a[i], a[pIndex] = a[pIndex], a[i]

		}
	}
	a[len(a)-1], a[pIndex+1] = a[pIndex+1], a[len(a)-1]

	return pIndex + 1
}
func main() {
	var ln int
	var err error
	var arr []int
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter lenght of array:")
	for {
		reader.Scan()
		text := reader.Text()
		ln, err = strconv.Atoi(text)
		if err != nil {
			fmt.Println("Enter integer only")
			continue
		}
		fmt.Println("")
		break
	}
	//generate array
	for i := 0; i < ln; i++ {
		arr = append(arr, rand.Intn(ln+100))
	}
	//sort.Ints(arr)

	tmp := make([]int, len(arr))

	copy(tmp, arr)

	start := time.Now()
	sortQuick(tmp)
	fmt.Println("\nQuick :", time.Since(start))

	copy(tmp, arr)

	start = time.Now()
	sortQuickRand(tmp)
	fmt.Println("\nQuickR:", time.Since(start))

	copy(tmp, arr)

	start = time.Now()
	sort.Ints(tmp)
	fmt.Println("\nGOLANGSort  :", time.Since(start))

	copy(tmp, arr)

	start = time.Now()

	tmp = sortmy.SimpleMergeSort(tmp)
	fmt.Println("\nSIMPLEMERGESort  :", time.Since(start))

	copy(tmp, arr)

	start = time.Now()

	tmp = sortmy.PARALLELMergeSort(tmp)
	fmt.Println("\nMultiMERGESort  :", time.Since(start))

	copy(tmp, arr)

	start = time.Now()

	tmp = sortmy.RunMMergeSort(tmp)
	fmt.Println("\nChannelMergeSort  :", time.Since(start))

	copy(tmp, arr)

	start = time.Now()

	tmp = sortmy.Insertionsort(tmp)
	fmt.Println("\nInsertionsort  :", time.Since(start))
}
