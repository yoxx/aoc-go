package main

import (
	"aoc-go/utils"
	"fmt"
	"regexp"
)

const airConst = "."
const sandConst = "o"
const rockConst = "#"

type Location struct {
	x            int
	y            int
	materialType string
}

func PartOne(inputStruct utils.FileStruct) int {
	//ParseInput(inputStruct)
	return 0
}

func PartTwo(inputStruct utils.FileStruct) int {
	// TODO write part two
	return 0
}

func ParseInput(inputStruct utils.FileStruct) (matrix map[int]map[int]Location) {
	matrix = make(map[int]map[int]Location)
	input, _ := utils.ReadFullFileInput(inputStruct)
	lines := utils.ParseLinesFromFullInput(input)
	// I want to fill my matrix but first need to determine my min/max X/Y to fill the grid
	for _, rockPath := range lines {
		rocks := regexp.MustCompile(" -> ").Split(rockPath, -1)
		if len(rocks) > 0 {
			for rockIndex, rock := range rocks {
				cords := regexp.MustCompile(",").Split(rock, -1)
				x := utils.MustParseStringToInt(cords[0])
				y := utils.MustParseStringToInt(cords[1])
				if matrix[y] == nil {
					matrix[y] = make(map[int]Location)
				}
				matrix[y][x] = Location{
					x:            x,
					y:            y,
					materialType: rockConst,
				}
				if rockIndex-1 >= 0 {
					prevCord := regexp.MustCompile(",").Split(rocks[rockIndex-1], -1)
					prevX := utils.MustParseStringToInt(prevCord[0])
					prevY := utils.MustParseStringToInt(prevCord[1])
					// Fill all the fields between the rocks aswell
					if prevX > x {
						amountOfRocksBetween := x - prevX
						for i := prevX - 1; prevX > amountOfRocksBetween; i-- {
							matrix[y][i] = Location{
								x:            i,
								y:            y,
								materialType: rockConst,
							}
						}
					} else if prevX < x {
						amountOfRocksBetween := x - prevX
						for i := prevX + 1; prevX < amountOfRocksBetween; i++ {
							matrix[y][i] = Location{
								x:            i,
								y:            y,
								materialType: rockConst,
							}
						}
					} else if prevY > y {
						amountOfRocksBetween := y - prevY
						for i := prevY - 1; prevY > amountOfRocksBetween; i-- {
							matrix[i][x] = Location{
								x:            x,
								y:            i,
								materialType: rockConst,
							}
						}
					} else if prevY < y {
						amountOfRocksBetween := y - prevY
						for i := prevY + 1; prevY < amountOfRocksBetween; i++ {
							matrix[i][x] = Location{
								x:            x,
								y:            i,
								materialType: rockConst,
							}
						}
					}
				}
			}
		}
	}
	return
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_14/day_14_input.txt"}
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
