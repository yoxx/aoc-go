package main

import (
	"aoc-go/utils"
	"fmt"
)

func PartOne(inputStruct utils.FileStruct) int {
	matrix := CreateMatrix(inputStruct)
	return DetermineVisibleTrees(matrix)
}

func PartTwo(inputStruct utils.FileStruct) int {
	matrix := CreateMatrix(inputStruct)
	return DetermineHighestTreeScenicScore(matrix)
}

func DetermineHighestTreeScenicScore(matrix [][]int) int {
	highestScore := 0
	for rIndex, row := range matrix {
		for cIndex, tree := range row {
			var leftScore, rightScore, upScore, downScore int
			// look left
			if cIndex == 0 {
				leftScore = 0
			} else {
				for i := cIndex - 1; i >= 0; i-- {
					// we always see a tree
					leftScore++
					if tree <= matrix[rIndex][i] {
						break
					}
				}
			}
			// look right
			if cIndex == len(row)-1 {
				rightScore = 0
			} else {
				for i := cIndex + 1; i <= len(row)-1; i++ {
					// we always see a tree
					rightScore++
					if tree <= matrix[rIndex][i] {
						break
					}
				}
			}
			// look up
			if rIndex == 0 {
				upScore = 0
			} else {
				for i := rIndex - 1; i >= 0; i-- {
					// we always see a tree
					upScore++
					if tree <= matrix[i][cIndex] {
						break
					}
				}
			}
			// look down
			if rIndex == len(matrix)-1 {
				downScore = 0
			} else {
				for i := rIndex + 1; i <= len(matrix)-1; i++ {
					// we always see a tree
					downScore++
					if tree <= matrix[i][cIndex] {
						break
					}
				}
			}
			var scenicScore = leftScore * rightScore * upScore * downScore
			if scenicScore > highestScore {
				highestScore = scenicScore
			}
		}
	}

	return highestScore
}

func DetermineVisibleTrees(matrix [][]int) (totalVisibleTrees int) {
	// Left
	leftMatrix := DetermineVisibleTreesOnSide(matrix)
	// Right
	rightMatrix := utils.ReverseMatrixSlices(matrix)
	rightMatrix = DetermineVisibleTreesOnSide(rightMatrix)
	rightMatrix = utils.ReverseMatrixSlices(rightMatrix)
	// Top
	topMatrix := utils.ReverseMatrixColumnsAndRows(matrix)
	topMatrix = DetermineVisibleTreesOnSide(topMatrix)
	topMatrix = utils.ReverseMatrixColumnsAndRows(topMatrix)
	// Bottom
	bottomMatrix := utils.ReverseMatrixColumnsAndRows(matrix)
	bottomMatrix = utils.ReverseMatrixSlices(bottomMatrix)
	bottomMatrix = DetermineVisibleTreesOnSide(bottomMatrix)
	bottomMatrix = utils.ReverseMatrixSlices(bottomMatrix)
	bottomMatrix = utils.ReverseMatrixColumnsAndRows(bottomMatrix)

	for rIndex, row := range rightMatrix {
		for cIndex, _ := range row {
			if leftMatrix[rIndex][cIndex] == 1 || rightMatrix[rIndex][cIndex] == 1 || topMatrix[rIndex][cIndex] == 1 || bottomMatrix[rIndex][cIndex] == 1 {
				totalVisibleTrees++
			}
		}
	}

	return totalVisibleTrees
}

func DetermineVisibleTreesOnSide(matrix [][]int) (visibleTreesMatrix [][]int) {
	// ALL edge trees are visible
	for rIndex, row := range matrix {
		visibleTreesMatrix = append(visibleTreesMatrix, []int{})
		highestTreeInRow := -1
		for _, col := range row {
			if col > highestTreeInRow {
				visibleTreesMatrix[rIndex] = append(visibleTreesMatrix[rIndex], 1)
				highestTreeInRow = col
			} else {
				visibleTreesMatrix[rIndex] = append(visibleTreesMatrix[rIndex], 0)
			}
		}
	}
	return visibleTreesMatrix
}

func CreateMatrix(inputStruct utils.FileStruct) (matrix [][]int) {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)
	for _, line := range lines {
		var row []int
		for _, char := range line {
			row = append(row, utils.MustParseStringToInt(string(char)))
		}
		if len(row) > 0 {
			matrix = append(matrix, row)
		}
	}
	return matrix
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_8/day_8_input.txt"}
	switch arguments.Part {
	case 1:
		fmt.Printf("Solution Part One: %d\n", PartOne(inputStruct))
		break
	case 2:
		fmt.Printf("Solution Part Two: %d\n", PartTwo(inputStruct))
	case 0:
		fmt.Printf("Solution Part One: %d\n", PartOne(inputStruct))
		fmt.Printf("Solution Part Two: %d\n", PartTwo(inputStruct))
	}
}
