package groupfinder

import (
	"test_task_yabbi/matrix"
	matrixgenerator "test_task_yabbi/matrixGenerator"
	"testing"
)

func benchmarkGroupFilter(height int, width int, colorsCount int, b *testing.B) {
	b.StopTimer()
	m := matrixgenerator.GenerateMatrixBySize(height, width, colorsCount)
	gf := New(m)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		gf.GetNumberOfElementsInLagrestGroup()
	}
}

func BenchmarkGroupFilter10(b *testing.B) {
	benchmarkGroupFilter(10, 10, 2, b)
}

func BenchmarkGroupFilter100(b *testing.B) {
	benchmarkGroupFilter(100, 100, 2, b)
}

func BenchmarkGroupFilter1000(b *testing.B) {
	benchmarkGroupFilter(1000, 1000, 2, b)
}

func BenchmarkGroupFilter5000(b *testing.B) {
	benchmarkGroupFilter(5000, 5000, 2, b)
}

func TestFindLargestGroup(t *testing.T) {
	m := matrix.Matrix{
		matrix.Row{1, 0, 1},
		matrix.Row{1, 0, 1},
		matrix.Row{1, 1, 1},
	}

	expected := 7
	groupFinder := New(m)

	actual := groupFinder.GetNumberOfElementsInLagrestGroup()

	if expected != actual {
		t.Fatal("Actual group size is incorrect")
	}

}