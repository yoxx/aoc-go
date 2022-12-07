package main

import (
	"aoc-go/utils"
	"fmt"
	"regexp"
)

type DirStruct struct {
	Parent       *DirStruct
	Name         string
	FileContents []FileStruct
	DirContents  map[string]*DirStruct // map[dirName]Dir
	TotalDirSize int
}

type FileStruct struct {
	Name string
	Size int
}

func PartOne(inputStruct utils.FileStruct) int {
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	rootDir := ParseInputToDirStruct(fullInput)
	_, sum, _ := DetermineDirSizeAndReturnSmallerThan(rootDir, 100000)
	return sum - rootDir.TotalDirSize
}

func PartTwo(inputStruct utils.FileStruct) int {
	maxUsedSpace := 70000000
	minUsedSpace := 30000000
	fullInput, _ := utils.ReadFullFileInput(inputStruct)
	rootDir := ParseInputToDirStruct(fullInput)
	_, _, currentUsedSpace := DetermineDirSizeAndReturnSmallerThan(rootDir, 100000)
	neededSpace := minUsedSpace - (maxUsedSpace - currentUsedSpace)
	dirsLargerThanMax, _ := DetermineDirSizeAndReturnLargerThanMaxSize(rootDir, neededSpace)

	var smallestDirMatchingCriteria *DirStruct
	for _, dir := range dirsLargerThanMax {
		if smallestDirMatchingCriteria == nil || dir.TotalDirSize < smallestDirMatchingCriteria.TotalDirSize {
			smallestDirMatchingCriteria = dir
		}
	}
	return smallestDirMatchingCriteria.TotalDirSize
}

func DetermineDirSizeAndReturnSmallerThan(dirStruct *DirStruct, sizeMax int) (dirSmallerThanSizeMax []*DirStruct, sumSmallerThenSizeMax int, totalSum int) {
	sumSmallerThenSizeMax = 0
	totalSum = 0
	// We need to determine the size for each directory (recursively)
	for _, fileContents := range dirStruct.FileContents {
		totalSum += fileContents.Size
	}
	for _, childDir := range dirStruct.DirContents {
		childDirSmallerThanSizeMax, childSmallerThenSizeMax, childSum := DetermineDirSizeAndReturnSmallerThan(childDir, sizeMax)
		dirSmallerThanSizeMax = append(dirSmallerThanSizeMax, childDirSmallerThanSizeMax...)
		totalSum += childSum
		sumSmallerThenSizeMax += childSmallerThenSizeMax
	}
	// We need to SUM the total sizes of these dirs (this CAN count files more than once as it bubbles-up)
	if totalSum < sizeMax {
		dirSmallerThanSizeMax = append(dirSmallerThanSizeMax, dirStruct)
		sumSmallerThenSizeMax += totalSum
	}

	dirStruct.TotalDirSize = totalSum
	return dirSmallerThanSizeMax, sumSmallerThenSizeMax, totalSum
}

func DetermineDirSizeAndReturnLargerThanMaxSize(dirStruct *DirStruct, sizeMax int) (dirLargerThanSizeMax []*DirStruct, totalSum int) {
	totalSum = 0
	// We need to determine the size for each directory (recursively)
	for _, fileContents := range dirStruct.FileContents {
		totalSum += fileContents.Size
	}
	for _, childDir := range dirStruct.DirContents {
		childDirLargerThanSizeMax, childSum := DetermineDirSizeAndReturnLargerThanMaxSize(childDir, sizeMax)
		dirLargerThanSizeMax = append(dirLargerThanSizeMax, childDirLargerThanSizeMax...)
		totalSum += childSum
	}
	// We need to SUM the total sizes of these dirs (this CAN count files more than once as it bubbles-up)
	if totalSum > sizeMax {
		dirLargerThanSizeMax = append(dirLargerThanSizeMax, dirStruct)
	}

	dirStruct.TotalDirSize = totalSum
	return dirLargerThanSizeMax, totalSum
}

func ParseInputToDirStruct(input []byte) *DirStruct {
	lines := utils.ParseLinesFromFullInput(input)
	// We have an input that consists of commands that move to/from or list directories and their contents
	// Loop through the lines and things starting with a $ are a command others are contents
	rootFolder := DirStruct{
		Name:        "/",
		DirContents: make(map[string]*DirStruct),
	}
	currentDir := &rootFolder
	for _, line := range lines {
		lineSlice := regexp.MustCompile(" ").Split(line, -1)
		if len(lineSlice) > 1 {
			if lineSlice[0] == "$" {
				// command
				switch lineSlice[1] {
				case "cd":
					if lineSlice[2] != "/" {
						// EITHER we go deeper or .. back
						if lineSlice[2] == ".." {
							currentDir = currentDir.Parent
						} else {
							// New dir
							currentDir = currentDir.DirContents[lineSlice[2]]
						}
					}
					break
				case "ls":
					// list contents :D ignore for now
				}
			} else {
				// contents
				if lineSlice[0] == "dir" {
					newDir := DirStruct{
						Parent:      currentDir,
						Name:        lineSlice[1],
						DirContents: make(map[string]*DirStruct),
					}
					currentDir.DirContents[lineSlice[1]] = &newDir
				} else {
					// file
					currentDir.FileContents = append(currentDir.FileContents, FileStruct{Name: lineSlice[1], Size: utils.MustParseStringToInt(lineSlice[0])})
				}
			}
		}
	}
	// we want to return the parent / uppermost / root folder
	return &rootFolder
}

func main() {
	arguments := utils.ParseCLIArguments()
	inputStruct := utils.FileStruct{Path: "./solutions/Y_2022/D_7/day_7_input.txt"}
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
