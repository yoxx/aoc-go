package main

import (
	"aoc-go/utils"
	"fmt"
	"sort"
)

func PartOne(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)

	currentElf := 0
	var elvesTotal []int
	elvesTotal = append(elvesTotal, 0)
	for _, line := range lines {
		if line == "" {
			elvesTotal = append(elvesTotal, 0)
			currentElf++
		} else {
			elvesTotal[currentElf] += utils.MustParseStringToInt(line)
		}
	}

	sort.Ints(elvesTotal)
	return elvesTotal[len(elvesTotal)-1]
}

func PartTwo(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)

	currentElf := 0
	var elvesTotal []int
	elvesTotal = append(elvesTotal, 0)
	for _, line := range lines {
		if line == "" {
			elvesTotal = append(elvesTotal, 0)
			currentElf++
		} else {
			elvesTotal[currentElf] += utils.MustParseStringToInt(line)
		}
	}

	sort.Ints(elvesTotal)
	return elvesTotal[len(elvesTotal)-1] + elvesTotal[len(elvesTotal)-2] + elvesTotal[len(elvesTotal)-3]
}

func main() {
	fmt.Println("Starting Puzzle 2022_1:")
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_1/day_1_input.txt"}
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
