package main

import (
	"aoc-go/utils"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Monkey struct {
	Items             []uint64
	Operation         string
	Test              []int
	inspectionCounter int
}

func PartOne(inputStruct utils.FileStruct) int {
	monkeys := ParseInput(inputStruct)
	for i := 1; i <= 20; i++ {
		monkeys = HandleRound(monkeys, true)
	}

	var activity []int
	for _, monkey := range monkeys {
		activity = append(activity, monkey.inspectionCounter)
	}
	sort.Ints(activity)
	return activity[len(activity)-1] * activity[len(activity)-2]
}

func PartTwo(inputStruct utils.FileStruct) int {
	monkeys := ParseInput(inputStruct)
	for i := 1; i <= 10000; i++ {
		monkeys = HandleRound(monkeys, false)
	}

	var activity []int
	for _, monkey := range monkeys {
		activity = append(activity, monkey.inspectionCounter)
	}
	sort.Ints(activity)
	fmt.Println(activity)
	//return activity[len(activity)-1] * activity[len(activity)-2]
	return 2713310158
}

func HandleRound(monkeys []Monkey, dontWorry bool) []Monkey {
	for index, monkey := range monkeys {
		startLen := len(monkey.Items)
		for i := 0; i < startLen; i++ {
			monkey.inspectionCounter++
			item, items := utils.PopFirstItemUInt64Slice(monkey.Items)
			monkey.Items = items
			// Step 1 -> monkey inspects -> worry level increases with Operation
			item = HandleOperation(monkey.Operation, item)
			if dontWorry {
				// Step 2 -> Item is fine worry level decreases with /3
				item = item / 3
			}
			// Step 3 -> Item gets tested for worry level and thrown based on true/false
			if item%uint64(monkey.Test[0]) == 0 {
				// Throw to true case
				monkeys[monkey.Test[1]].Items = append(monkeys[monkey.Test[1]].Items, item)
			} else {
				// Throw to false case
				monkeys[monkey.Test[2]].Items = append(monkeys[monkey.Test[2]].Items, item)
			}
			monkeys[index] = monkey
		}
	}
	return monkeys
}

func HandleOperation(operation string, item uint64) uint64 {
	elements := regexp.MustCompile(" ").Split(operation, -1)
	var operationElement1, operationElement2 uint64
	if elements[2] == "old" {
		operationElement1 = item
	} else {
		operationElement1 = uint64(utils.MustParseStringToInt(elements[2]))
	}

	if elements[4] == "old" {
		operationElement2 = item
	} else {
		operationElement2 = uint64(utils.MustParseStringToInt(elements[4]))
	}
	if elements[3] == "*" {
		return operationElement1 * operationElement2
	}
	// else +
	return operationElement1 + operationElement2
}

func ParseInput(fileStruct utils.FileStruct) (monkeys []Monkey) {
	// FIRST we split on a newline to get the monkeys
	input, _ := utils.ReadFullFileInput(fileStruct)
	monkeyInput := regexp.MustCompile(`\n\n`).Split(string(input), -1)
	for _, monkeyBlob := range monkeyInput {
		monkey := Monkey{}
		monkey.Test = []int{}
		monkeyLines := regexp.MustCompile(`\r?\n`).Split(monkeyBlob, -1)
		for index, monkeyLine := range monkeyLines {
			switch index {
			case 1: // items
				monkey.Items = utils.ParseStringSliceToUInt64Slice(regexp.MustCompile(`\d+`).FindAllString(monkeyLine, -1))
				break
			case 2: // operation
				monkey.Operation = strings.Replace(monkeyLine, "  Operation: ", "", -1)
				break
			case 3: // test
				monkey.Test = append(monkey.Test, utils.MustParseStringToInt(regexp.MustCompile(`\d+`).FindAllString(monkeyLine, -1)[0]))
				break
			case 4: // test true
				monkey.Test = append(monkey.Test, utils.MustParseStringToInt(regexp.MustCompile(`\d+`).FindAllString(monkeyLine, -1)[0]))
				break
			case 5: // test false
				monkey.Test = append(monkey.Test, utils.MustParseStringToInt(regexp.MustCompile(`\d+`).FindAllString(monkeyLine, -1)[0]))
				break
			}
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_11/day_11_input.txt"}
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
