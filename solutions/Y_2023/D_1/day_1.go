package main

import (
	"aoc-go/utils"
	"fmt"
	"github.com/dlclark/regexp2"
	"regexp"
)

func PartOne(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)

	totalSum := 0
	for _, line := range lines {
		re := regexp.MustCompile("([0-9])")
		numbers := re.FindAllString(line, -1)

		// if numbers is empty, continue
		if len(numbers) != 0 {
			charFirstResult := []rune(numbers[0])
			charLastResult := []rune(numbers[len(numbers)-1])
			currentSum := string(charFirstResult[0:1]) + string(charLastResult[len(charLastResult)-1:])
			//fmt.Printf("charFirstResult: %s, charLastResult: %s, currentSum: %s\n", string(charFirstResult), string(charLastResult), currentSum)

			totalSum += utils.MustParseStringToInt(currentSum)
		}
	}
	return totalSum
}

func PartTwo(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(fullInput)

	totalSum := 0
	for _, line := range lines {
		re := regexp2.MustCompile("(?=(zero|one|two|three|four|five|six|seven|eight|nine))|([0-9])", 0)

		numbers := regexp2FindAllString(re, line)

		// if numbers is empty, continue
		if len(numbers) != 0 {
			charFirstResult := []rune(parseNumber(numbers[0]))
			charLastResult := []rune(parseNumber(numbers[len(numbers)-1]))
			currentSum := string(charFirstResult[0:1]) + string(charLastResult[len(charLastResult)-1:])
			//fmt.Printf("line: %s, charFirstResult: %s, charLastResult: %s, currentSum: %s\n", line, string(charFirstResult), string(charLastResult), currentSum)

			totalSum += utils.MustParseStringToInt(currentSum)
		}
	}
	return totalSum
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2023/D_1/day_1_input.txt"}
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

func parseNumber(number string) string {
	switch number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	case "zero":
		return "0"
	default:
		return number
	}
}

func regexp2FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		groups := m.Groups()
		for _, group := range groups {
			if len(group.Captures) > 0 && group.Captures[0].Length > 0 {
				matches = append(matches, group.Captures[0].String())
			}
		}
		m, _ = re.FindNextMatch(m)
	}
	return matches
}
