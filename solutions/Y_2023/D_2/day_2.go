package main

import (
	"aoc-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func PartOne(inputStruct utils.FileStruct) int {
	input, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(input)

	sum := 0
	for _, line := range lines {
		if line != "" {
			id, possible := playGame(line)
			if possible {
				sum += id
			}
		}
	}

	return sum
}

func PartTwo(inputStruct utils.FileStruct) int {
	input, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(input)

	sum := 0
	for _, line := range lines {
		if line != "" {
			_, power := playGame2(line)
			// echo the value of power
			fmt.Printf("Power: %d\n", power)
			sum += power
		}
	}

	return sum
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2023/D_2/day_2_input.txt"}
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

func playGame(line string) (id int, possible bool) {
	possible = true
	id, input := parseLine(line)

	for _, gameInput := range input {
		for _, red := range gameInput.red {
			if red > 12 {
				return id, false
			}
		}
		for _, green := range gameInput.green {
			if green > 13 {
				return id, false
			}
		}
		for _, blue := range gameInput.blue {
			if blue > 14 {
				return id, false
			}
		}

	}

	return id, possible
}

func playGame2(line string) (id int, power int) {
	id, input := parseLine(line)

	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, gameInput := range input {
		for _, red := range gameInput.red {
			if red > maxRed {
				maxRed = red
			}
		}
		for _, green := range gameInput.green {
			if green > maxGreen {
				maxGreen = green
			}
		}
		for _, blue := range gameInput.blue {
			if blue > maxBlue {
				maxBlue = blue
			}
		}

	}

	return id, maxRed * maxGreen * maxBlue
}

func parseLine(line string) (id int, input []gameInput) {
	parts := strings.Split(line, ":")
	id, _ = strconv.Atoi(strings.Split(parts[0], " ")[1])

	input = []gameInput{}
	draws := strings.Split(parts[1], ";")
	for _, draw := range draws {
		gameInput := gameInput{}
		parts := strings.Split(draw, ",")
		for _, part := range parts {
			comp := strings.Split(strings.Trim(part, " "), " ")

			i, _ := strconv.Atoi(comp[0])
			if strings.Trim(comp[1], " ") == "red" {
				gameInput.red = append(gameInput.red, i)
				continue
			}

			if strings.Trim(comp[1], " ") == "blue" {
				gameInput.blue = append(gameInput.blue, i)
				continue
			}

			if strings.Trim(comp[1], " ") == "green" {
				gameInput.green = append(gameInput.green, i)
			}
		}
		input = append(input, gameInput)
	}

	return id, input
}

type gameInput struct {
	red   []int
	green []int
	blue  []int
}
