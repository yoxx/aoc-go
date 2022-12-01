package main

import (
	"aoc-go/utils"
	"fmt"
	"regexp"
)

func PartOne(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)
	validCount := 0
	for _, line := range lines {
		regex := regexp.MustCompile(`\d+`)
		nums := regex.FindAllStringSubmatch(line, -1)
		if len(nums) > 0 {
			if utils.MustParseStringToInt(nums[0][0])+utils.MustParseStringToInt(nums[1][0]) > utils.MustParseStringToInt(nums[2][0]) &&
				utils.MustParseStringToInt(nums[1][0])+utils.MustParseStringToInt(nums[2][0]) > utils.MustParseStringToInt(nums[0][0]) &&
				utils.MustParseStringToInt(nums[2][0])+utils.MustParseStringToInt(nums[0][0]) > utils.MustParseStringToInt(nums[1][0]) {
				validCount++
			}
		}
	}
	return validCount
}

func PartTwo(inputStruct utils.FileStruct) int {
	// TODO write part two
	return 0
}

func main() {
	fmt.Println("Starting Puzzle 2016_3:")
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2016/D_3/day_3_input.txt"}
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
