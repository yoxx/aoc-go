package main

import (
	"aoc-go/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func PartOne(inputStruct utils.FileStruct) int {
	// get the contents of the file and split it into two lists of numbers
	fullFile, _ := utils.ReadFullFileInput(inputStruct)
	contents := utils.ParseLinesFromFullInput(fullFile)
	// split the contents []string into two lists of numbers
	var list1, list2 []int
	for _, line := range contents {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			num1, _ := strconv.Atoi(parts[0])
			num2, _ := strconv.Atoi(parts[1])
			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
	}
	// sort the lists from lowest to highest
	sort.Ints(list1)
	sort.Ints(list2)
	// take the difference between both lowests on the lists and add up the differences
	var sum int
	for i := 0; i < len(list1); i++ {
		// abs the integer difference between the two numbers
		sum += int(math.Abs(float64(list2[i] - list1[i])))
	}

	return sum
}

func PartTwo(inputStruct utils.FileStruct) int {
	// get the contents of the file and split it into two lists of numbers
	fullFile, _ := utils.ReadFullFileInput(inputStruct)
	contents := utils.ParseLinesFromFullInput(fullFile)
	// split the contents []string into two lists of numbers
	var list1, list2 []int
	for _, line := range contents {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			num1, _ := strconv.Atoi(parts[0])
			num2, _ := strconv.Atoi(parts[1])
			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
	}
	// sort the lists from lowest to highest
	sort.Ints(list1)
	sort.Ints(list2)
	// take the difference between both lowests on the lists and add up the differences
	var similarity int
	for i := 0; i < len(list1); i++ {
		var amountOfTimesInRightList int
		for i2 := 0; i2 < len(list2); i2++ {
			if list1[i] == list2[i2] {
				amountOfTimesInRightList++
			}
		}
		similarity += list1[i] * amountOfTimesInRightList
	}

	return similarity
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2024/D_1/day_1_input.txt"}
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
