package main

import (
	"aoc-go/utils"
	"fmt"
)

func PartOne(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	return FindStartOfPacketMarker([]rune(string(fullInput)))
}

func PartTwo(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	return FindStartOfMessageMarker([]rune(string(fullInput)))
}

func FindStartOfPacketMarker(signal []rune) int {
	var output int
	for index, _ := range signal {
		var checkSlice []rune
		var signalChunk []rune
		for i := 0; i < 4; i++ {
			signalChunk = append(signalChunk, signal[index+i])
		}
		checkSlice = utils.RuneIntersectionUnique(signalChunk, signalChunk)
		if len(checkSlice) == 4 {
			output = index + 4
			break
		}
	}
	return output
}

func FindStartOfMessageMarker(signal []rune) int {

	var output int
	for index, _ := range signal {
		var checkSlice []rune
		var signalChunk []rune
		for i := 0; i < 14; i++ {
			signalChunk = append(signalChunk, signal[index+i])
		}
		checkSlice = utils.RuneIntersectionUnique(signalChunk, signalChunk)
		if len(checkSlice) == 14 {
			output = index + 14
			break
		}
	}
	return output
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_6/day_6_input.txt"}
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
