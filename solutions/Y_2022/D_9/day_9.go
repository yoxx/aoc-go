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
		if len(lineSlice) > 1 {
			newHeadLocations, newTailLocations := MoveRopeHeadAndTail(headLocations[len(headLocations)-1], tailLocations[len(tailLocations)-1], lineSlice[0], utils.MustParseStringToInt(lineSlice[1]))
			tailLocations = append(tailLocations, newTailLocations...)
			headLocations = append(headLocations, newHeadLocations...)
		}
	}
	return len(DedupLocationSlice(tailLocations))
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
			curHeadLoc.y--
			if !InSquareAround(lastTailLoc, curHeadLoc) {
				curTailLoc = MoveTailCloserToHead(curTailLoc, curHeadLoc)
			}
			break
		case "D":
			curHeadLoc.y++
			if !InSquareAround(lastTailLoc, curHeadLoc) {
				curTailLoc = MoveTailCloserToHead(curTailLoc, curHeadLoc)
			}
			break
		case "R":
			curHeadLoc.x++
			if !InSquareAround(lastTailLoc, curHeadLoc) {
				curTailLoc = MoveTailCloserToHead(curTailLoc, curHeadLoc)
			}
			break
		case "L":
			curHeadLoc.x--
			if !InSquareAround(lastTailLoc, curHeadLoc) {
				curTailLoc = MoveTailCloserToHead(curTailLoc, curHeadLoc)
			}
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

func MoveTailCloserToHead(curTail location, curHead location) location {
	// Down
	if curHead.x == curTail.x && curHead.y > curTail.y {
		curTail.y++
	}
	// Up
	if curHead.x == curTail.x && curHead.y < curTail.y {
		curTail.y--
	}
	// Left
	if curHead.x < curTail.x && curHead.y == curTail.y {
		curTail.x--
	}
	// Right
	if curHead.x > curTail.x && curHead.y == curTail.y {
		curTail.x++
	}
	// Diagonal Left Up
	if curHead.x < curTail.x && curHead.y < curTail.y {
		curTail.y--
		curTail.x--
	}
	// Diagonal Right Up
	if curHead.x > curTail.x && curHead.y < curTail.y {
		curTail.y--
		curTail.x++
	}
	// Diagonal Left Down
	if curHead.x < curTail.x && curHead.y > curTail.y {
		curTail.y++
		curTail.x--
	}
	// Diagonal Right Down
	if curHead.x > curTail.x && curHead.y > curTail.y {
		curTail.y++
		curTail.x++
	}

	return curTail
}

func InSquareAround(tailLocation location, headLocation location) bool {
	if headLocation.x >= tailLocation.x-1 && headLocation.x <= tailLocation.x+1 &&
		headLocation.y >= tailLocation.y-1 && headLocation.y <= tailLocation.y+1 {
		return true
	}
	return false
}

func DedupLocationSlice(locations []location) []location {
	var intersection []location
	charMap := make(map[location]bool, len(locations))
	// Setup a map
	for _, item := range locations {
		charMap[item] = true
	}
	// IF we have the same char append as intersection
	for _, item := range locations {
		if charMap[item] {
			// Dedup on the GO
			if !LocationSliceContains(intersection, item) {
				intersection = append(intersection, item)
			}
		}
	}
	return intersection
}

func LocationSliceContains(slice []location, item location) bool {
	for _, value := range slice {
		if value.x == item.x && value.y == item.y {
			return true
		}
	}
	return false
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
