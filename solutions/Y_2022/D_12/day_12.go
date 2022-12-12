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
	paths := FindPaths(endIndex, startIndex, heightMap, []Location{})
	// Get the shortest path
	shortestPath := 0
	for _, path := range paths {
		if shortestPath == 0 || len(path) < shortestPath {
			shortestPath = len(path)
		}
	}
	//return shortestPath
	return shortestPath - 1
}

func PartTwo(inputStruct utils.FileStruct) int {
	// TODO write part two
	return 0
}

func FindPaths(startIndex Location, endIndex Location, heightMap [][]Location, pathSoFar []Location) (paths [][]Location) {
	// Right-io lets walk all the paths we can find
	// We are only allowed to go Right, Left, Up, Down
	// We are only allowed to visit places that are of the same elevation OR 1 lower/higher
	// For ease of use we are putting the locations we visited to visited = true
	// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
	pathSoFar = append(pathSoFar, heightMap[startIndex.y][startIndex.x])
	fmt.Println(string(startIndex.char), startIndex)
	// Much wow we reached the end lets save this path
	if startIndex == endIndex {
		return append(paths, pathSoFar)
	}
	// Up
	if startIndex.y-1 >= 0 && !isValueInList(heightMap[startIndex.y-1][startIndex.x], pathSoFar) && heightMap[startIndex.y-1][startIndex.x].char >= (startIndex.char-1) && heightMap[startIndex.y-1][startIndex.x].char <= (startIndex.char+1) {
		path := FindPaths(heightMap[startIndex.y-1][startIndex.x], endIndex, heightMap, pathSoFar)
		if len(path) > 0 {
			paths = append(paths, path...)
		}
	}
	// Down
	if startIndex.y+1 < len(heightMap) && !isValueInList(heightMap[startIndex.y+1][startIndex.x], pathSoFar) && heightMap[startIndex.y+1][startIndex.x].char >= (startIndex.char-1) && heightMap[startIndex.y+1][startIndex.x].char <= (startIndex.char+1) {
		path := FindPaths(heightMap[startIndex.y+1][startIndex.x], endIndex, heightMap, pathSoFar)
		if len(path) > 0 {
			paths = append(paths, path...)
		}
	}
	// Left
	if startIndex.x-1 >= 0 && !isValueInList(heightMap[startIndex.y][startIndex.x-1], pathSoFar) && heightMap[startIndex.y][startIndex.x-1].char >= (startIndex.char-1) && heightMap[startIndex.y][startIndex.x-1].char <= (startIndex.char+1) {
		path := FindPaths(heightMap[startIndex.y][startIndex.x-1], endIndex, heightMap, pathSoFar)
		if len(path) > 0 {
			paths = append(paths, path...)
		}
	}
	// Right
	if startIndex.x+1 < len(heightMap[0]) && !isValueInList(heightMap[startIndex.y][startIndex.x+1], pathSoFar) && heightMap[startIndex.y][startIndex.x+1].char >= (startIndex.char-1) && heightMap[startIndex.y][startIndex.x+1].char <= (startIndex.char+1) {
		path := FindPaths(heightMap[startIndex.y][startIndex.x+1], endIndex, heightMap, pathSoFar)
		if len(path) > 0 {
			paths = append(paths, path...)
		}
	}
	return
}

func InputToHeightMap(inputStruct utils.FileStruct) (startIndex Location, endIndex Location, heightMap [][]Location) {
	input, _ := utils.ReadFullFileInput(inputStruct)
	for yIndex, line := range utils.ParseLinesFromFullInput(input) {
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
