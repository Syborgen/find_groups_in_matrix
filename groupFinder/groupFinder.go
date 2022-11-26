package groupfinder

import (
	"test_task_yabbi/matrix"
)

type groupFinder struct {
	ColorMatrix     matrix.Matrix
	GroupMatrix     matrix.Matrix
	ElementsInGroup map[int]int
	LastGroupNumber int
}

func New(m matrix.Matrix) groupFinder {
	return groupFinder{
		ColorMatrix:     m,
		GroupMatrix:     matrix.New(len(m), len(m[0])),
		ElementsInGroup: make(map[int]int),
	}
}

func (gf *groupFinder) GetLagrestGroup() (int, int) {
	gf.fillGroupMatrix()

	max := 0
	maxGroup := 0
	for group, count := range gf.ElementsInGroup {
		if count > max {
			max = count
			maxGroup = group
		}
	}

	return maxGroup, max
}

func (gf *groupFinder) fillGroupMatrix() {
	for i, row := range gf.ColorMatrix {
		for j := range row {
			groupNumber, elementsCount := gf.tryMarkGroup(i, j)
			if elementsCount == 0 {
				continue
			}

			gf.ElementsInGroup[groupNumber] = elementsCount
		}
	}
}

func (gf *groupFinder) tryMarkGroup(i int, j int) (int, int) {
	groupNumber := gf.GroupMatrix[i][j]

	if groupNumber != 0 {
		return groupNumber, 0
	}

	gf.LastGroupNumber++
	groupElementsCount := gf.recursiveMarkGroup(i, j)
	gf.ElementsInGroup[gf.LastGroupNumber] = groupElementsCount

	return gf.GroupMatrix[i][j], groupElementsCount
}

func (gf *groupFinder) recursiveMarkGroup(i int, j int) int {
	gf.GroupMatrix[i][j] = gf.LastGroupNumber

	result := 1

	if gf.isRightSameGroupAndUnmarked(i, j) {
		result += gf.recursiveMarkGroup(i, j+1)
	}

	if gf.isBotSameGroupAndUnmarked(i, j) {
		result += gf.recursiveMarkGroup(i+1, j)
	}

	if gf.isLeftSameGroupAndUnmarked(i, j) {
		result += gf.recursiveMarkGroup(i, j-1)
	}

	if gf.isTopSameGroupAndUnmarked(i, j) {
		result += gf.recursiveMarkGroup(i-1, j)
	}

	return result
}

func (gf *groupFinder) isRightSameGroupAndUnmarked(i int, j int) bool {
	maxWidth := len(gf.ColorMatrix[0]) - 1
	return j < maxWidth && gf.isRightSameColor(i, j) && !gf.isMarked(i, j+1)
}

func (gf *groupFinder) isRightSameColor(i int, j int) bool {
	return gf.ColorMatrix[i][j] == gf.ColorMatrix[i][j+1]
}

func (gf *groupFinder) isBotSameGroupAndUnmarked(i int, j int) bool {
	maxHeight := len(gf.ColorMatrix) - 1
	return i < maxHeight && gf.isBotSameColor(i, j) && !gf.isMarked(i+1, j)
}

func (gf *groupFinder) isLeftSameGroupAndUnmarked(i int, j int) bool {
	return j > 0 && gf.isLeftSameColor(i, j) && !gf.isMarked(i, j-1)
}

func (gf *groupFinder) isLeftSameColor(i int, j int) bool {
	return gf.ColorMatrix[i][j] == gf.ColorMatrix[i][j-1]
}

func (gf *groupFinder) isTopSameGroupAndUnmarked(i int, j int) bool {
	return i > 0 && gf.isTopSameColor(i, j) && !gf.isMarked(i-1, j)
}

func (gf *groupFinder) isTopSameColor(i int, j int) bool {
	return gf.ColorMatrix[i][j] == gf.ColorMatrix[i-1][j]
}

func (gf *groupFinder) isBotSameColor(i int, j int) bool {
	return gf.ColorMatrix[i][j] == gf.ColorMatrix[i+1][j]
}

func (gf *groupFinder) isMarked(i int, j int) bool {
	return gf.GroupMatrix[i][j] != 0
}
