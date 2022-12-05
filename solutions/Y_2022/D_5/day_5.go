package main

import (
	"aoc-go/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func PartOne(inputStruct utils.FileStruct) string {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	stack, procedures := parseStackAndCommands(fullInput)

	for _, procedureLine := range procedures {
		matches := regexp.MustCompile(`\d+`).FindAllString(procedureLine, -1)
		if len(matches) > 0 {
			for numberOfCrateToMove := 0; numberOfCrateToMove < utils.MustParseStringToInt(matches[0]); numberOfCrateToMove++ {
				// Stack match[1] to match[2]
				poppedVal := stack[utils.MustParseStringToInt(matches[1])-1][len(stack[utils.MustParseStringToInt(matches[1])-1])-1]
				stack[utils.MustParseStringToInt(matches[1])-1] = stack[utils.MustParseStringToInt(matches[1])-1][:len(stack[utils.MustParseStringToInt(matches[1])-1])-1]
				stack[utils.MustParseStringToInt(matches[2])-1] = append(stack[utils.MustParseStringToInt(matches[2])-1], poppedVal)
			}
		}
	}

	var output string
	for _, stackCol := range stack {
		output += stackCol[len(stackCol)-1]
	}
	return output
}

func PartTwo(inputStruct utils.FileStruct) string {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	stack, procedures := parseStackAndCommands(fullInput)

	for _, procedureLine := range procedures {
		matches := regexp.MustCompile(`\d+`).FindAllString(procedureLine, -1)
		if len(matches) > 0 {
			var movingStack []string
			for numberOfCrateToMove := 0; numberOfCrateToMove < utils.MustParseStringToInt(matches[0]); numberOfCrateToMove++ {
				// Stack match[1] to match[2]
				poppedVal := stack[utils.MustParseStringToInt(matches[1])-1][len(stack[utils.MustParseStringToInt(matches[1])-1])-1]
				stack[utils.MustParseStringToInt(matches[1])-1] = stack[utils.MustParseStringToInt(matches[1])-1][:len(stack[utils.MustParseStringToInt(matches[1])-1])-1]
				movingStack = append(movingStack, poppedVal)
			}
			movingStack = utils.ReverseStringSlice(movingStack)
			for _, crate := range movingStack {
				stack[utils.MustParseStringToInt(matches[2])-1] = append(stack[utils.MustParseStringToInt(matches[2])-1], crate)
			}
		}
	}

	var output string
	for _, stackCol := range stack {
		output += stackCol[len(stackCol)-1]
	}
	return output
}

func parseStackAndCommands(input []byte) ([][]string, []string) {
	contents := regexp.MustCompile(`\n\n`).Split(string(input), -1)

	var cargoStack [][]string

	stackLines := utils.ParseLinesFromFullInput([]byte(contents[0]))
	rows := int(math.Ceil(float64(len([]byte(stackLines[0]))) / 4))
	for i := 0; i < rows; i++ {
		// Fill our slices
		cargoStack = append(cargoStack, []string(nil))
	}
	for index, stackLine := range stackLines {
		if index != len(stackLines)-1 {
			line := []rune(stackLine)
			for i := 0; i < len(line); i += 4 {
				var crateLine []rune
				if len(line[i:]) < 4 {
					crateLine = line[i:]
				} else {
					crateLine = line[i : i+4]
				}
				crate := strings.Trim(string(crateLine), "\n []")
				if len(crate) > 0 && rune(crate[0]) > 0 {
					cargoStack[i/4] = append(cargoStack[i/4], crate)
				}
			}
		}
	}
	for i := 0; i < rows; i++ {
		// flip our slices
		cargoStack[i] = utils.ReverseStringSlice(cargoStack[i])
	}

	procedure := utils.ParseLinesFromFullInput([]byte(contents[1]))

	return cargoStack, procedure
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_5/day_5_input.txt"}
	switch arguments.Part {
	case 1:
		fmt.Printf("Solution Part One: %s\n", PartOne(inputStruct))
		break
	case 2:
		fmt.Printf("Solution Part Two: %s\n", PartTwo(inputStruct))
	case 0:
		fmt.Printf("Solution Part One: %s\n", PartOne(inputStruct))
		fmt.Printf("Solution Part Two: %s\n", PartTwo(inputStruct))
	}
}
