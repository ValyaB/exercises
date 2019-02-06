package main

import "fmt"

func play(n int, s, t, d *[]int) {
	fmt.Printf("Source: %v, Temp: %v, Dest: %v N: %v\n", s, t, d, n)
	if n == 1 {
		*d = append(*d, (*s)[len(*s)-1])
		*s = (*s)[:len(*s)-1]
		i++
		fmt.Printf("-Source: %v, Temp: %v, Dest: %v\n", s, t, d)
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
	source := []int{4, 3, 2, 1}
	dest, temp := []int{}, []int{}
	i = 0
	play(len(source), &source, &temp, &dest)
	fmt.Println(i)
}
