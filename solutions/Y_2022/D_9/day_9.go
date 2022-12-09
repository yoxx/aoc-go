package main

import (
	"aoc-go/utils"
	"fmt"
	"os"
	"regexp"
)

type location struct {
	x int
	y int
}

func PartOne(inputStruct utils.FileStruct) int {
	tailLocations := []location{{0, 0}}
	headLocations := []location{{0, 0}}
	input, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(input)
	for _, line := range lines {
		lineSlice := regexp.MustCompile(" ").Split(line, -1)
		newHeadLocations, newTailLocations := MoveRopeHeadAndTail(headLocations[len(headLocations)-1], headLocations[len(headLocations)-1], lineSlice[0], utils.MustParseStringToInt(lineSlice[1]))
		println(newHeadLocations, newTailLocations)
		tailLocations = append(tailLocations, newTailLocations...)
		tailLocations = append(headLocations, newHeadLocations...)
	}
	return 0
}

func PartTwo(inputStruct utils.FileStruct) int {
	// TODO write part two
	return 0
}

func MoveRopeHeadAndTail(curHeadLoc location, curTailLoc location, direction string, steps int) (headLocations []location, tailLocations []location) {
	lastTailLoc := location{curTailLoc.x, curTailLoc.y}
	for step := 0; step < steps; step++ {
		switch direction {
		case "U":
			break
		case "D":
			break
		case "R":
			break
		case "L":
			break
		default:
			fmt.Printf("UNKNOWN DIRECTION CASE: %s", direction)
			os.Exit(1)
		}
		// The head always moves
		headLocations = append(headLocations, curHeadLoc)
		if lastTailLoc.x != curTailLoc.x || lastTailLoc.y != curTailLoc.y {
			// Tail has moved!
			tailLocations = append(tailLocations, curTailLoc)
			lastTailLoc = location{curTailLoc.x, curTailLoc.y}
		}
	}
	return headLocations, tailLocations
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_9/day_9_input.txt"}
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
