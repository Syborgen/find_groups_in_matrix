package matrix

import (
	"strings"
)

type Matrix []Row

func New(height int, width int) Matrix {
	matrix := make(Matrix, height)
	for i := range matrix {
		matrix[i] = make(Row, width)
	}

	return matrix
}

func (m Matrix) String() string {
	matrixAsStringSlice := make([]string, 0, len(m))
	for _, row := range m {
		matrixAsStringSlice = append(matrixAsStringSlice, row.Stirng())
	}

	return strings.Join(matrixAsStringSlice, "\n")
}
