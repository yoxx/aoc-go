package main

import (
	"aoc-go/utils"
	"fmt"
)

type Location struct {
	char rune
	x    int
	y    int
}

func PartOne(inputStruct utils.FileStruct) int {
	startIndex, endIndex, heightMap := InputToHeightMap(inputStruct)
	return BFS(endIndex, startIndex, heightMap)
}

func PartTwo(inputStruct utils.FileStruct) int {
	// TODO write part two
	return 0
}

func BFS(startLocation, finish Location, heightMap [][]Location) int {
	yMax := len(heightMap)
	xMax := len(heightMap[0])
	// https://www.youtube.com/watch?v=xlVX7dXLS64
	visited := map[Location]bool{}
	queue := []Location{startLocation}
	distanceToLoc := map[Location]int{startLocation: 0}
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		// Early out if we reached our target
		poppedLocation := queue[0]
		if poppedLocation == finish {
			break
		}
		queue = queue[1:]
		if !visited[poppedLocation] {
			visited[poppedLocation] = true
			// Get the neighbours
			for _, dir := range directions {
				nextX := poppedLocation.x + dir[0]
				nextY := poppedLocation.y + dir[1]

				// Early out for non existing (out-of-bounds) locations
				if nextX < 0 || nextX >= xMax || nextY < 0 || nextY >= yMax {
					continue
				}
				// Get our next location
				nextLocation := Location{
					char: heightMap[poppedLocation.y+dir[1]][poppedLocation.x+dir[0]].char,
					x:    poppedLocation.x + dir[0],
					y:    poppedLocation.y + dir[1],
				}
				// Check if we already visited them
				if !visited[nextLocation] {
					if nextLocation.char == poppedLocation.char || nextLocation.char == poppedLocation.char-1 {
						queue = append(queue, nextLocation)
						if distanceToLoc[nextLocation] == 0 {
							// We have not yet calculated a distance to this location
							distanceToLoc[nextLocation] = distanceToLoc[poppedLocation] + 1
						} else if distanceToLoc[nextLocation] >= distanceToLoc[poppedLocation]+1 {
							distanceToLoc[nextLocation] = distanceToLoc[poppedLocation] + 1
						}
					}
				}
			}
		}
	}
	return distanceToLoc[finish]
}

func InputToHeightMap(inputStruct utils.FileStruct) (startIndex Location, endIndex Location, heightMap [][]Location) {
	input, _ := utils.ReadFullFileInput(inputStruct)
	for yIndex, line := range utils.ParseLinesFromFullInput(input) {
		if len(line) > 0 {
			heightMap = append(heightMap, []Location{})
			for xIndex, char := range line {
				location := Location{
					char: char,
					x:    xIndex,
					y:    yIndex,
				}
				if char == 'S' {
					location.char = 'a'
					startIndex = location
				}
				if char == 'E' {
					location.char = 'z'
					endIndex = location
				}
				heightMap[yIndex] = append(heightMap[yIndex], location)
			}
		}
	}
	return
}

func isValueInList(value Location, list []Location) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_12/day_12_input.txt"}
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
