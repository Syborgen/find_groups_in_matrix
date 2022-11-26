package matrixgenerator

import (
	"math/rand"
	"test_task_yabbi/matrix"
)

func GenerateMatrixBySize(height int, width int,
	colorsCount int) matrix.Matrix {
	var matrix matrix.Matrix = make([]matrix.Row, 0, height)
	for i := 0; i < height; i++ {
		newRow := generateRow(width, colorsCount)
		matrix = append(matrix, newRow)
	}

	return matrix
}

func generateRow(width int, colorsCount int) matrix.Row {
	var result matrix.Row = make([]int, 0, width)
	for i := 0; i < width; i++ {
		result = append(result, rand.Intn(colorsCount))
	}

	return result
}
