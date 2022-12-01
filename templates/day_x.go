package main

import (
	"aoc-go/utils"
	"fmt"
)

func PartOne(inputStruct utils.FileStruct) int {
	fmt.Println("YX_DX Running Part One")
	// TODO write part two
	return 0
}

func PartTwo(inputStruct utils.FileStruct) int {
	fmt.Println("YX_DX Running Part Two")
	// TODO write part two
	return 0
}

func main() {
	fmt.Println("Starting Puzzle YX_DX:")
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_YX/D_DX/day_DX_input.txt"}
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
