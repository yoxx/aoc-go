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
	// TODO write part two
	return 0
}

//func DetermineHighestTreeScenicScore(matrix [][]int) int {
//	highestScore := 0
//	for rIndex, row := range matrix {
//		for cIndex, col := row {
//			if (rIndex == 0 || )
//		}
//	}
//}

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
		matrix = append(matrix, row)
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
