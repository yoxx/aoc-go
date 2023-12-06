package main

import (
	"aoc-go/utils"
	"fmt"
	"regexp"
)

func PartOne(inputStruct utils.FileStruct) int {
	fileinput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fileinput)

	regex := regexp.MustCompile(`[\d]+`)
	timeList := regex.FindAllString(lines[0], -1)
	distanceList := regex.FindAllString(lines[1], -1)

	totalWaysToBeatRecord := 1
	for i := 0; i < len(timeList); i++ {
		time := utils.MustParseStringToInt(timeList[i])
		distance := utils.MustParseStringToInt(distanceList[i])

		waysToBeatRecord := 0
		for t := 0; t < time; t++ {
			speed := t * 1
			timeLeft := time - t
			distancePossible := speed * timeLeft

			if distancePossible > distance {
				waysToBeatRecord += 1
			}
		}

		totalWaysToBeatRecord *= waysToBeatRecord
	}
	return totalWaysToBeatRecord
}

func PartTwo(inputStruct utils.FileStruct) int {
	fileinput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fileinput)

	regex := regexp.MustCompile(`[\d]+`)
	timeList := regex.FindAllString(lines[0], -1)
	distanceList := regex.FindAllString(lines[1], -1)

	totalTime := ""
	for i := 0; i < len(timeList); i++ {
		totalTime += timeList[i]
	}
	totalDistance := ""
	for i := 0; i < len(distanceList); i++ {
		totalDistance += distanceList[i]
	}

	time := utils.MustParseStringToInt(totalTime)
	distance := utils.MustParseStringToInt(totalDistance)

	waysToBeatRecord := 0
	for t := 0; t < time; t++ {
		speed := t * 1
		timeLeft := time - t
		distancePossible := speed * timeLeft

		if distancePossible > distance {
			waysToBeatRecord += 1
		}
	}

	return waysToBeatRecord
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2023/D_6/day_6_input.txt"}
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
