package main

import (
	"aoc-go/utils"
	"fmt"
	"regexp"
	"unicode"
)

type keys struct {
	startAt int
	endsAt  int
}

type parts struct {
	number int
	row    int
	keys   keys
}

func PartOne(inputStruct utils.FileStruct) int {
	input, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(input)

	partNumbers := []parts{}
	// Loop through the lines array.
	for key, line := range lines {
		var prevLine []rune
		// check if lines[key-1] exists
		if key > 0 {
			prevLine = []rune(lines[key-1])
		}
		var nextLine []rune
		if key+1 < len(lines) {
			nextLine = []rune(lines[key+1])
		}
		curLine := []rune(line)
		for i, char := range curLine {
			if unicode.IsNumber(char) {
				// check for every rune the surrounding runes, even diagonally
				// Previous line
				if len(prevLine) > 0 {
					if i > 0 {
						if isSymbol(prevLine[i-1]) {
							start, end, number := findNumberInLine(curLine, i)
							partNumbers = addToPartList(partNumbers, parts{number: number, keys: keys{startAt: start, endsAt: end}})
						}
					}
					if isSymbol(prevLine[i]) {
						start, end, number := findNumberInLine(curLine, i)
						partNumbers = addToPartList(partNumbers, parts{number: number, keys: keys{startAt: start, endsAt: end}})
					}
					if i+1 < len(prevLine) && isSymbol(prevLine[i+1]) {
						start, end, number := findNumberInLine(curLine, i)
						partNumbers = addToPartList(partNumbers, parts{number: number, keys: keys{startAt: start, endsAt: end}})
					}
				}
				// Current line
				if i > 0 {
					if isSymbol(curLine[i-1]) {
						start, end, number := findNumberInLine(curLine, i)
						partNumbers = addToPartList(partNumbers, parts{number: number, keys: keys{startAt: start, endsAt: end}})
					}
				}
				if i+1 < len(curLine) && isSymbol(curLine[i+1]) {
					start, end, number := findNumberInLine(curLine, i)
					partNumbers = addToPartList(partNumbers, parts{number: number, keys: keys{startAt: start, endsAt: end}})
				}
				if len(nextLine) > 0 {
					if i > 0 {
						if isSymbol(nextLine[i-1]) {
							start, end, number := findNumberInLine(curLine, i)
							partNumbers = addToPartList(partNumbers, parts{number: number, keys: keys{startAt: start, endsAt: end}})
						}
					}
					if isSymbol(nextLine[i]) {
						start, end, number := findNumberInLine(curLine, i)
						partNumbers = addToPartList(partNumbers, parts{number: number, keys: keys{startAt: start, endsAt: end}})
					}
					if i+1 < len(nextLine) && isSymbol(nextLine[i+1]) {
						start, end, number := findNumberInLine(curLine, i)
						partNumbers = addToPartList(partNumbers, parts{number: number, keys: keys{startAt: start, endsAt: end}})
					}
				}
			}
		}
	}

	// loop through the partlist and sum all the numbers
	sum := 0
	for _, part := range partNumbers {
		sum += part.number
	}
	return sum
}

func PartTwo(inputStruct utils.FileStruct) int {
	// TODO write part two
	return 0
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2023/D_3/day_3_input.txt"}
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

func addToPartList(partNumbers []parts, partNumber parts) []parts {
	for _, part := range partNumbers {
		if part.number == partNumber.number && part.row == partNumber.row && part.keys.startAt == partNumber.keys.startAt && part.keys.endsAt == partNumber.keys.endsAt {
			return partNumbers
		}
	}
	partNumbers = append(partNumbers, partNumber)
	return partNumbers
}

func findNumberInLine(runelist []rune, key int) (start int, end int, number int) {
	rStr := string(runelist)
	r := regexp.MustCompile("[0-9]+")
	indexes := r.FindAllIndex([]byte(rStr), -1)
	matches := r.FindAllString(string(runelist), -1)
	for i, index := range indexes {
		if key >= index[0] && key <= index[1] {
			return index[0], index[1], utils.MustParseStringToInt(matches[i])
		}
	}
	return 0, 0, 0
}

func isSymbol(r rune) bool {
	return r != 46 && (r < 48 || r > 57)
}
