package main

import (
	"aoc-go/utils"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

func PartOne(inputStruct utils.FileStruct) int {
	indicesSum := 0
	inputPairs := ParseInput(inputStruct)
	for index, pair := range inputPairs {
		if ApplyRules(pair[0], pair[1]) <= 0 {
			indicesSum += index + 1
		}
	}
	return indicesSum
}

func PartTwo(inputStruct utils.FileStruct) int {
	// TODO write part two
	return 0
}

func ApplyRules(firstPair, secondPair any) int {
	firstVal, fCheck := firstPair.(float64)
	secondVal, sCheck := secondPair.(float64)

	// IF we have two integers (they are ok) then we are done
	if fCheck && sCheck {
		// Check if the second number us higher the result should be 0 or -n
		return int(firstVal) - int(secondVal)
	}
	// It were not numbers slices then? Or a number and a slice / slice and a number
	var firstValSlice, secondValSlice []any
	switch t := firstPair.(type) {
	case []any, []float64: // If we have a slice or a slice of floats pass it on
		firstValSlice = t.([]any)
		break
	case float64: // If we have a number (then the other side must be a slice) pass it on as slice
		firstValSlice = []any{t}
		break
	}
	switch t := secondPair.(type) {
	case []any, []float64: // If we have a slice or a slice of floats pass it on
		secondValSlice = t.([]any)
		break
	case float64: // If we have a number (then the other side must be a slice) pass it on as slice
		secondValSlice = []any{t}
		break
	}

	// Loop through the slices and get them index by index for each slice (first and second)
	for index := range firstValSlice {
		// IF the length of our second slice is smaller its no good
		if len(secondValSlice) <= index {
			return 1
		}
		result := ApplyRules(firstValSlice[index], secondValSlice[index])
		// IF our result is 0 we have the same number and need to continue to check the next number
		if result != 0 {
			return result
		}
	}
	if len(firstValSlice) == len(secondValSlice) {
		// Our slices have the same length continue
		return 0
	}
	return -1
}

func ParseInput(inputStruct utils.FileStruct) (pairs [][]any) {
	input, _ := utils.ReadFullFileInput(inputStruct)
	var inputPairs [][]string
	for _, blob := range regexp.MustCompile("\n\n").Split(string(input), -1) {
		inputPairs = append(inputPairs, regexp.MustCompile("\n").Split(blob, -1))
	}

	for _, inputPair := range inputPairs {
		var firstPair any
		var secondPair any
		if err := json.Unmarshal([]byte(inputPair[0]), &firstPair); err != nil {
			fmt.Printf("Cannot unmarshal this string into JSON: %s", inputPair[0])
			os.Exit(1)
		}

		if err := json.Unmarshal([]byte(inputPair[1]), &secondPair); err != nil {
			fmt.Printf("Cannot unmarshal this string into JSON: %s", inputPair[1])
			os.Exit(1)
		}
		var pair []any
		pair = append(pair, firstPair, secondPair)
		pairs = append(pairs, pair)
	}
	return
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_13/day_13_input.txt"}
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
