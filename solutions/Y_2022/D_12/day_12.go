package main

import (
	"aoc-go/utils"
	"fmt"
)

type Location struct {
	char    rune
	x       int
	y       int
	visited bool
}

func PartOne(inputStruct utils.FileStruct) int {
	startIndex, endIndex, heigthMap := InputToHeightMap(inputStruct)
	paths := FindPaths(endIndex, startIndex, heigthMap, []Location{})
	// Get the shortest path
	shortestPath := 0
	for _, path := range paths {
		if shortestPath == 0 || len(path) < shortestPath {
			shortestPath = len(path)
		}
	}
	//return shortestPath
	return 0
}

func PartTwo(inputStruct utils.FileStruct) int {
	// TODO write part two
	return 0
}

func FindPaths(startIndex Location, endIndex Location, heigthMap [][]Location, pathsSoFar []Location) (paths [][]Location) {
	// Right-io lets walk all the paths we can find
	// We are only allowed to go Right, Left, Up, Down
	// We are only allowed to visit places that are of the same elevation OR 1 lower/higher
	// For ease of use we are putting the locations we visited to visited = true
	// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
	// Much wow we reached the end lets save this path
	if startIndex.x == endIndex.x && startIndex.y == endIndex.y {
		paths = append(paths, pathsSoFar)
	}
	if startIndex.char == 'S' {
		startIndex.char = 'a'
	}
	// Up
	if startIndex.y-1 != 0 && !heigthMap[startIndex.y-1][startIndex.x].visited && heigthMap[startIndex.y-1][startIndex.x].char >= (startIndex.char-1) && heigthMap[startIndex.y-1][startIndex.x].char <= (startIndex.char+1) {
		pathsSoFar = append(pathsSoFar, heigthMap[startIndex.y-1][startIndex.x])
		paths = FindPaths(heigthMap[startIndex.y-1][startIndex.x], endIndex, heigthMap, pathsSoFar)
	}
	// Down
	if startIndex.y+1 < len(heigthMap) && !heigthMap[startIndex.y+1][startIndex.x].visited && heigthMap[startIndex.y+1][startIndex.x].char >= (startIndex.char-1) && heigthMap[startIndex.y+1][startIndex.x].char <= (startIndex.char+1) {
		pathsSoFar = append(pathsSoFar, heigthMap[startIndex.y+1][startIndex.x])
		paths = FindPaths(heigthMap[startIndex.y+1][startIndex.x], endIndex, heigthMap, pathsSoFar)
	}
	// Left
	if startIndex.y+1 < len(heigthMap) && !heigthMap[startIndex.y+1][startIndex.x].visited && heigthMap[startIndex.y+1][startIndex.x].char >= (startIndex.char-1) && heigthMap[startIndex.y+1][startIndex.x].char <= (startIndex.char+1) {
		pathsSoFar = append(pathsSoFar, heigthMap[startIndex.y+1][startIndex.x])
		paths = FindPaths(heigthMap[startIndex.y+1][startIndex.x], endIndex, heigthMap, pathsSoFar)
	}
	// Right
	return
}

func InputToHeightMap(inputStruct utils.FileStruct) (startIndex Location, endIndex Location, heightMap [][]Location) {
	input, _ := utils.ReadFullFileInput(inputStruct)
	for yIndex, line := range utils.ParseLinesFromFullInput(input) {
		heightMap = append(heightMap, []Location{})
		for xIndex, char := range line {
			location := Location{
				char:    char,
				x:       xIndex,
				y:       yIndex,
				visited: false,
			}
			heightMap[yIndex] = append(heightMap[yIndex], location)
			if char == 'S' {
				startIndex = location
			}
			if char == 'E' {
				endIndex = location
			}
		}
	}
	return
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
