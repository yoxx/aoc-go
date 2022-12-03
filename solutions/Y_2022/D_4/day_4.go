package main

import (
	"aoc-go/utils"
	"fmt"
	"regexp"
)

func PartOne(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)
	overlappingRanges := 0
	for _, line := range lines {
		ranges := regexp.MustCompile(`(\d+)`).FindAllString(line, -1)
		if len(ranges) > 0 {
			if (utils.MustParseStringToInt(ranges[0]) >= utils.MustParseStringToInt(ranges[2]) && utils.MustParseStringToInt(ranges[1]) <= utils.MustParseStringToInt(ranges[3])) ||
				(utils.MustParseStringToInt(ranges[2]) >= utils.MustParseStringToInt(ranges[0]) && utils.MustParseStringToInt(ranges[3]) <= utils.MustParseStringToInt(ranges[1])) {
				overlappingRanges++
			}
		}
	}
	return overlappingRanges
}

func PartTwo(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)
	overlappingRanges := 0
	for _, line := range lines {
		ranges := regexp.MustCompile(`(\d+)`).FindAllString(line, -1)
		if len(ranges) > 0 {
			if (utils.MustParseStringToInt(ranges[0]) >= utils.MustParseStringToInt(ranges[2]) && utils.MustParseStringToInt(ranges[0]) <= utils.MustParseStringToInt(ranges[3])) ||
				(utils.MustParseStringToInt(ranges[1]) >= utils.MustParseStringToInt(ranges[2]) && utils.MustParseStringToInt(ranges[1]) <= utils.MustParseStringToInt(ranges[3])) ||
				(utils.MustParseStringToInt(ranges[2]) >= utils.MustParseStringToInt(ranges[0]) && utils.MustParseStringToInt(ranges[2]) <= utils.MustParseStringToInt(ranges[1])) ||
				(utils.MustParseStringToInt(ranges[3]) >= utils.MustParseStringToInt(ranges[0]) && utils.MustParseStringToInt(ranges[3]) <= utils.MustParseStringToInt(ranges[1])) {
				overlappingRanges++
			}
		}
	}
	return overlappingRanges
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_4/day_4_input.txt"}
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
