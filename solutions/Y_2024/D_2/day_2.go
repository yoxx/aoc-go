package main

import (
	"aoc-go/utils"
	"fmt"
	"math"
	"strings"
)

func PartOne(inputStruct utils.FileStruct) int {
	// get the contents of the file and split it into two lists of numbers
	fullFile, _ := utils.ReadFullFileInput(inputStruct)
	contents := utils.ParseLinesFromFullInput(fullFile)
	safeReports := 0
	for _, line := range contents {
		lineIntSlice := utils.ParseStringSliceToIntSlice(strings.Fields(line))

		if len(lineIntSlice) > 0 && validateReport(lineIntSlice, false) {
			safeReports++
		}
	}
	return safeReports
}

func PartTwo(inputStruct utils.FileStruct) int {
	// get the contents of the file and split it into two lists of numbers
	fullFile, _ := utils.ReadFullFileInput(inputStruct)
	contents := utils.ParseLinesFromFullInput(fullFile)
	safeReports := 0
	lineCounter := 0
	for _, line := range contents {
		lineCounter++
		lineIntSlice := utils.ParseStringSliceToIntSlice(strings.Fields(line))

		if len(lineIntSlice) > 0 && validateReport(lineIntSlice, true) {
			safeReports++
		}
	}
	return safeReports
}

func validateReport(lineIntSlice []int, problemDampner bool) bool {
	increase := false
	decrease := false
	for i := 0; i < len(lineIntSlice); i++ {
		if i+1 != len(lineIntSlice) {
			diff := math.Abs(float64(lineIntSlice[i] - lineIntSlice[i+1]))
			if diff == 0 || diff > 3 {
				if problemDampner {
					// remove the offending (next) number and recheck
					copySlice := make([]int, len(lineIntSlice))
					copy(copySlice, lineIntSlice)
					newSlice := append(copySlice[:i+1], copySlice[i+2:]...)
					if validateReport(newSlice, false) {
						return true
					} else {
						// try and removing the current number
						copySlice2 := make([]int, len(lineIntSlice))
						copy(copySlice2, lineIntSlice)
						newSlice2 := append(copySlice2[:i], copySlice2[i+1:]...)
						return validateReport(newSlice2, false)
					}
				}
				return false
			}
			if lineIntSlice[i] > lineIntSlice[i+1] && !increase {
				decrease = true
			} else if lineIntSlice[i] < lineIntSlice[i+1] && !decrease {
				increase = true
			} else {
				if problemDampner {
					// remove the offending (next) number and recheck
					copySlice := make([]int, len(lineIntSlice))
					copy(copySlice, lineIntSlice)
					newSlice := append(copySlice[:i+1], copySlice[i+2:]...)
					if validateReport(newSlice, false) {
						return true
					} else {
						// try and removing the current number
						copySlice2 := make([]int, len(lineIntSlice))
						copy(copySlice2, lineIntSlice)
						newSlice2 := append(copySlice2[:i], copySlice2[i+1:]...)
						if validateReport(newSlice2, false) {
							return true
						} else {
							// try and removing the FIRST number
							copySlice3 := make([]int, len(lineIntSlice))
							copy(copySlice3, lineIntSlice)
							newSlice3 := copySlice3[1:]
							return validateReport(newSlice3, false)
						}
					}
				}
				return false
			}
		}
	}

	return true
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2024/D_2/day_2_input.txt"}
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
