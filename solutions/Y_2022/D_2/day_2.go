package main

import (
	"aoc-go/utils"
	"fmt"
	"regexp"
)

func PartOne(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)

	total := 0
	for _, line := range lines {
		regex := regexp.MustCompile(`[A-Z]`)
		matches := regex.FindAllString(line, -1)
		if len(matches) > 0 {
			total += rockPaperScissorsPart1(matches[0], matches[1])
		}
	}
	return total
}

func PartTwo(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)

	total := 0
	for _, line := range lines {
		regex := regexp.MustCompile(`[A-Z]`)
		matches := regex.FindAllString(line, -1)
		if len(matches) > 0 {
			total += rockPaperScissorsPart2(matches[0], matches[1])
		}
	}
	return total
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_2/day_2_input.txt"}
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

func rockPaperScissorsPart1(opponent string, myself string) int {
	total := 0
	switch opponent {
	case "A": // Rock
		switch myself {
		case "X": //ROCK DRAW
			total += 3
			break
		case "Y": //PAPER WON
			total += 6
			break
		case "Z": //SCISSORS LOST
			total += 0
			break
		}
		break
	case "B": // Paper
		switch myself {
		case "X": //ROCK LOST
			total += 0
			break
		case "Y": //PAPER DRAW
			total += 3
			break
		case "Z": //SCISSORS WIN
			total += 6
			break
		}
		break
	case "C": // Scissors
		switch myself {
		case "X": //ROCK WIN
			total += 6
			break
		case "Y": //PAPER LOST
			total += 0
			break
		case "Z": //SCISSORS DRAW
			total += 3
			break
		}
		break
	}
	return total + determineWorth(myself)
}

func rockPaperScissorsPart2(opponent string, myself string) int {
	total := 0
	switch opponent {
	case "A": // Rock
		switch myself {
		case "X": // LOSE
			total += 0 + determineWorth("Z") // Scissors
			break
		case "Y": // DRAW
			total += 3 + determineWorth("X") // Rock
			break
		case "Z": // WIN
			total += 6 + determineWorth("Y") // Paper
			break
		}
		break
	case "B": // Paper
		switch myself {
		case "X": // LOSE
			total += 0 + determineWorth("X") // Rock
			break
		case "Y": // DRAW
			total += 3 + determineWorth("Y") // Paper
			break
		case "Z": // WIN
			total += 6 + determineWorth("Z") // Scissors
			break
		}
		break
	case "C": // Scissors
		switch myself {
		case "X": // LOSE
			total += 0 + determineWorth("Y") // Paper
			break
		case "Y": // DRAW
			total += 3 + determineWorth("Z") // Scissors
			break
		case "Z": // WIN
			total += 6 + determineWorth("X") // Rock
			break
		}
		break
	}
	return total
}

func determineWorth(shape string) int {
	worth := 0
	switch shape {
	case "X":
		worth = 1
		break
	case "Y":
		worth = 2
		break
	case "Z":
		worth = 3
		break
	}
	return worth
}
