package main

import (
	"fmt"
	groupfinder "test_task_yabbi/groupFinder"
	matrixgenerator "test_task_yabbi/matrixGenerator"
)

func main() {
	matrix := matrixgenerator.GenerateMatrixBySize(3, 3, 2)

	fmt.Println("Generated matrix")
	fmt.Println(matrix)

	gf := groupfinder.New(matrix)
	max := gf.GetNumberOfElementsInLagrestGroup()

	fmt.Println("Group matrix")
	fmt.Println(gf.GroupMatrix)

	fmt.Println("Elements in group (groupNumber:elementsInGroup)")
	fmt.Println(gf.ElementsInGroup)

	fmt.Println("Max group size:", max)
}
