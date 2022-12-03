package main

import (
	"aoc-go/utils"
	"fmt"
)

func PartOne(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)

	totalSum := 0

	for _, line := range lines {
		// Each rugsack has 2 compartments which is the string split in the middle
		firstCompartment := []rune(line[:len(line)/2])
		secondCompartment := []rune(line[len(line)/2:])
		// Every char is a different item but there can be multiple of the same items

		intersect := utils.RuneIntersectionUnique(firstCompartment, secondCompartment)
		if len(intersect) > 0 {
			totalSum += MapValueToPrio(intersect[0])
		}
	}
	return totalSum
}

func PartTwo(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)

	totalSum := 0
	var groupSlices [][]rune
	var intersect []rune
	for _, line := range lines {
		groupSlices = append(groupSlices, []rune(line))
		if len(groupSlices) >= 3 {
			intersect = utils.RuneIntersectionUnique(groupSlices[0], groupSlices[1])
			intersect = utils.RuneIntersectionUnique(intersect, groupSlices[2])
			totalSum += MapValueToPrio(intersect[0])
			intersect = nil
			groupSlices = nil
		}
	}
	return totalSum
}

func MapValueToPrio(item rune) int {
	prioMap := make(map[rune]int, 52)

	var asciiStartLowercase = 'a'
	var asciiStartUppercase = 'A'
	for i := 1; i <= 52; i++ {
		if asciiStartLowercase <= 'z' {
			prioMap[asciiStartLowercase] = i
			asciiStartLowercase++
		} else {
			prioMap[asciiStartUppercase] = i
			asciiStartUppercase++
		}
	}

	return prioMap[item]
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_3/day_3_input.txt"}
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
