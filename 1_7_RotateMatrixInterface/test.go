package exercises

func rotateRow(layers ...interface{}) []interface{} {

	for i := range layers {
		layers[i], layers[0] = layers[0], layers[i]
	}
	return layers
}

func transposeMatrix(matrix [][]interface{}) [][]interface{} {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				continue
			}
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}
func verticalReflactionMatrix(matrix [][]interface{}) [][]interface{} {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}

	}
	return matrix
}
func horizontalReflactionMatrix(matrix [][]interface{}) [][]interface{} {
	n := len(matrix)
	for i := 0; i < n/2; i++ {

		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]

	}
	return matrix
}

func rotateMatrix90(matrix [][]interface{}) [][]interface{} {
	return horizontalReflactionMatrix(transposeMatrix(matrix))
}
