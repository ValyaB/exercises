package main

import (
	"fmt"
	"strconv"
)

var checked = make(map[string]bool)
var count int

type mymatrix [][]string

func (m *mymatrix) checkSameConnectedElements(i, j, localmax int) {
	if len(*m) == 0 || len(*m) == 1 && len((*m)[i]) == 0 {
		return
	}
	if _, ok := checked[strconv.Itoa(i)+strconv.Itoa(j)]; !ok {
		checked[strconv.Itoa(i)+strconv.Itoa(j)] = true
		localmax++

		if i < len(*m)-1 {
			if (*m)[i+1][j] == (*m)[i][j] {

				m.checkSameConnectedElements(i+1, j, localmax)
			}
		}
		if i > 0 {
			if (*m)[i-1][j] == (*m)[i][j] {

				m.checkSameConnectedElements(i-1, j, localmax)
			}
		}
		if j < len((*m)[i])-1 {
			if (*m)[i][j+1] == (*m)[i][j] {
				m.checkSameConnectedElements(i, j+1, localmax)
			}
		}
		if j > 0 {
			if (*m)[i][j-1] == (*m)[i][j] {
				m.checkSameConnectedElements(i, j-1, localmax)
			}
		}

		if count < localmax {
			count = localmax
			fmt.Println("full end", checked, "count", count, (*m)[i][j])
		}

		if j < len((*m)[i])-1 {
			m.checkSameConnectedElements(i, j+1, 0)
		} else {
			if i < len((*m))-1 {
				m.checkSameConnectedElements(i+1, 0, 0)
			}
		}
	}
}
func main() {

	var m mymatrix
	// m = [][]string{{"c", "b", "c"}, {"c", "k", "c"}, {"c", "c", "c"}}
	// for i := range m {
	// 	fmt.Println(m[i])
	// }
	// m.checkSameConnectedElements(0, 0, 0)
	// count = 0
	// m = nil

	// checked = nil
	m = [][]string{{"ac", "c", "bc"}, {"c", "k", "c"}, {"c", "kc", "c"}}

	for i := range m {
		fmt.Println(m[i])
	}
	m.checkSameConnectedElements(0, 0, 0)
}
