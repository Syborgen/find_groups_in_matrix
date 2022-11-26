package main

import (
	"fmt"
	groupfinder "test_task_yabbi/groupFinder"
	matrixgenerator "test_task_yabbi/matrixGenerator"
)

func main() {
	matrix := matrixgenerator.GenerateMatrixBySize(5, 5, 3)

	fmt.Println("Generated matrix")
	fmt.Println(matrix)

	gf := groupfinder.New(matrix)
	maxGroup, maxGroupSize := gf.GetLagrestGroup()

	fmt.Println("Matrix with the largest group")
	gf.PrintMatrixWithMarkedGroup(maxGroup)

	fmt.Printf("Max group size %d\n", maxGroupSize)
}
