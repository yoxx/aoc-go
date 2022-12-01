package main

import (
	"aoc-go/utils"
	"fmt"
)

func main() {
	arguments := utils.ParseCLIArguments()
	switch arguments.Part {
	case 1:
		fmt.Printf("Solution Part One: %d\n", PartOne())
		break
	case 2:
		fmt.Printf("Solution Part Two: %d\n", PartTwo())
	case 0:
		fmt.Printf("Solution Part One: %d\n", PartOne())
		fmt.Printf("Solution Part Two: %d\n", PartTwo())
	}
}

func PartOne() int {
	fmt.Println("Running Part One")
	// TODO write part two
	return 0
}

func PartTwo() int {
	fmt.Println("Running Part Two")
	// TODO write part two
	return 0
}
