package exercises

import "fmt"

func zeroMatrix(matrix [][]interface{}) [][]interface{} {
	rows, columns := []int{}, []int{}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				columns = append(columns, j)
				rows = append(rows, i)
				break
			}
		}
	}
	fmt.Println(rows, columns)
	for _, r := range rows {
		for j := 0; j < len(matrix[r]); j++ {
			matrix[r][j] = 0
		}
	}
	for _, c := range columns {
		for i := 0; i < len(matrix); i++ {
			matrix[i][c] = 0
		}
	}
	return matrix
}
