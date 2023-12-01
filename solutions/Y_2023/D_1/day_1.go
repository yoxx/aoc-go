package main

import (
	"aoc-go/utils"
	"fmt"
	"regexp"
)

func PartOne(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)

	totalSum := 0
	for _, line := range lines {
		re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(line, -1)

		// if numbers is empty, continue
		if len(numbers) != 0 {
			charFirstResult := []rune(numbers[0])
			charLastResult := []rune(numbers[len(numbers)-1])
			currentSum := string(charFirstResult[0:1]) + string(charLastResult[len(charLastResult)-1:])
			//fmt.Printf("charFirstResult: %s, charLastResult: %s, currentSum: %s\n", string(charFirstResult), string(charLastResult), currentSum)

			totalSum += utils.MustParseStringToInt(currentSum)
		}
	}
	return totalSum
}

func PartTwo(inputStruct utils.FileStruct) int {
	// TODO write part two
	return 0
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2023/D_1/day_1_input.txt"}
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
