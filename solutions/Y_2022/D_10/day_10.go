package main

import (
	"aoc-go/utils"
	"fmt"
	"regexp"
)

func PartOne(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)

	var subSums []int
	cycleCounter := 0
	valueX := 1
	regex := regexp.MustCompile(" ")
	for _, line := range lines {
		command := regex.Split(line, -1)
		if len(command) > 0 {
			switch command[0] {
			case "noop":
				cycleCounter++
				subSums = CheckAndSaveSignalStrenghtIfNeeded(subSums, cycleCounter, valueX)
				break
			case "addx":
				cycleCounter++
				subSums = CheckAndSaveSignalStrenghtIfNeeded(subSums, cycleCounter, valueX)
				cycleCounter++
				subSums = CheckAndSaveSignalStrenghtIfNeeded(subSums, cycleCounter, valueX)
				valueX += utils.MustParseStringToInt(command[1])
			}
		}

		if cycleCounter > 220 {
			break
		}
	}
	totalSum := 0
	for _, value := range subSums {
		totalSum += value
	}

	return totalSum
}

func PartTwo(inputStruct utils.FileStruct) (output string) {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)

	cycleCounter := 0
	valueX := 1
	regex := regexp.MustCompile(" ")
	for _, line := range lines {
		command := regex.Split(line, -1)
		if len(command) > 0 {
			switch command[0] {
			case "noop":
				cycleCounter++
				output += DrawCRT(cycleCounter, valueX)
				break
			case "addx":
				cycleCounter++
				output += DrawCRT(cycleCounter, valueX)
				cycleCounter++
				output += DrawCRT(cycleCounter, valueX)
				valueX += utils.MustParseStringToInt(command[1])
			}
		}
	}

	return output
}

func DrawCRT(currentCycle int, xPos int) (output string) {
	currentCTRPosition := currentCycle%40 - 1
	if currentCTRPosition == -1 {
		currentCTRPosition = 39
	}
	if currentCTRPosition >= xPos-1 && currentCTRPosition <= xPos+1 {
		output += "#"
	} else {
		output += "."
	}
	if currentCTRPosition == 39 {
		output += "\n"
	}
	//fmt.Printf(output)

	return output
}

func CheckAndSaveSignalStrenghtIfNeeded(subvalues []int, counter int, value int) []int {
	switch counter {
	case 20:
		subvalues = append(subvalues, value*20)
		break
	case 60:
		subvalues = append(subvalues, value*60)
		break
	case 100:
		subvalues = append(subvalues, value*100)
		break
	case 140:
		subvalues = append(subvalues, value*140)
		break
	case 180:
		subvalues = append(subvalues, value*180)
		break
	case 220:
		subvalues = append(subvalues, value*220)
		break
	}
	return subvalues
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_10/day_10_input.txt"}
	switch arguments.Part {
	case 1:
		fmt.Printf("Solution Part One: %d\n", PartOne(inputStruct))
		break
	case 2:
		fmt.Printf("Solution Part Two: \n%s", PartTwo(inputStruct))
	case 0:
		fmt.Printf("Solution Part One: %d\n", PartOne(inputStruct))
		fmt.Printf("Solution Part Two: \n%s", PartTwo(inputStruct))
	}
}
