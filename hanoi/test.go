package main

import "fmt"

func play(n int, s, t, d *[]int) {
	//fmt.Printf("Source: %v, Temp: %v, Dest: %v N: %v\n", s, t, d, n)
	if n == 1 {
		*d = append(*d, (*s)[len(*s)-1])
		*s = (*s)[:len(*s)-1]
		i++
		//	fmt.Printf("-Source: %v, Temp: %v, Dest: %v\n", s, t, d)
		return
	}

	play(n-1, s, d, t)

	*d = append(*d, (*s)[len(*s)-1])
	*s = (*s)[:len(*s)-1]
	i++

	play(n-1, t, s, d)

}

var i int

func main() {
	source := []int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70, 69, 68, 67, 66, 65, 64, 63, 62, 61, 60, 4, 3, 2, 1}
	dest, temp := []int{}, []int{}
	i = 0
	play(len(source), &source, &temp, &dest)
	fmt.Println(i)
}
