package utils

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/cavaliergopher/grab/v3"
	"github.com/joho/godotenv"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

type FileStruct struct {
	Path     string
	Contents []byte
}

type CliArguments struct {
	Command string
	Year    int
	Day     int
	Part    int
}

func ParseCLIArguments() (arguments CliArguments) {
	var command string
	var year int
	var day int
	var part int
	flag.StringVar(&command, "command", "default", "The command to run, possible options are:\n"+
		"\t* puzzle - run a puzzle of the provided Year / Day / Part\n"+
		"\t* input - Download the input of the given Year / Day"+
		"\t* test - Download the input of the given Year / Day")
	flag.IntVar(&year, "year", 0, "The year you would like to download input or run the puzzles from")
	flag.IntVar(&day, "day", 0, "The day you would like to download input or run the puzzles from")
	flag.IntVar(&part, "part", 0, "The part of the day you would like to run / test")

	flag.Parse()
	return CliArguments{
		Command: command,
		Year:    year,
		Day:     day,
		Part:    part,
	}
}

func RunPuzzle(arguments CliArguments) {
	fmt.Printf("Running puzzle Y%d_D%d\n", arguments.Year, arguments.Day)
	path := fmt.Sprintf("./solutions/Y_%d/D_%d/day_%d.go", arguments.Year, arguments.Day, arguments.Day)
	cmd := exec.Command("go", "run", path, "-part", fmt.Sprintf("%d", arguments.Part))
	out, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", out)

	if err != nil {
		fmt.Printf("Encountered errors while running puzzle Y%d_D%d: %v\n", arguments.Year, arguments.Day, err)
		os.Exit(1)
	}
}

func TestPuzzle(arguments CliArguments) {
	path := fmt.Sprintf("./solutions/Y_%d/D_%d", arguments.Year, arguments.Day)

	switch arguments.Part {
	case 1:
		fmt.Printf("Running TestPartOne for puzzle Y%d_D%d\n", arguments.Year, arguments.Day)
		cmd := exec.Command("go", "test", "-run", "TestPartOne", path, "-v")
		out, err := cmd.CombinedOutput()
		fmt.Printf("%s\n", out)
		if err != nil {
			fmt.Printf("Encountered errors while running TestPartOne for puzzle Y%d_D%d: %v\n", arguments.Year, arguments.Day, err)
			os.Exit(1)
		}
		break
	case 2:
		fmt.Printf("Running TestPartTwo for puzzle Y%d_D%d\n", arguments.Year, arguments.Day)
		cmd := exec.Command("go", "test", "-run", "TestPartTwo", path, "-v")
		out, err := cmd.CombinedOutput()
		fmt.Printf("%s\n", out)
		if err != nil {
			fmt.Printf("Encountered errors while running TestPartTwo for puzzle Y%d_D%d: %v\n", arguments.Year, arguments.Day, err)
			os.Exit(1)
		}
		break
	case 0:
		fmt.Printf("Running ALL unittests for puzzle Y%d_D%d\n", arguments.Year, arguments.Day)
		cmd := exec.Command("go", "test", path, "-v")
		out, err := cmd.CombinedOutput()
		fmt.Printf("%s\n", out)
		if err != nil {
			fmt.Printf("Encountered errors while running all unittests for puzzle Y%d_D%d: %v\n", arguments.Year, arguments.Day, err)
			os.Exit(1)
		}
	}
}

func DownloadInput(arguments CliArguments) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		fmt.Println("Error loading .env, was this file created?")
		os.Exit(1)
	}

	// Validate the year input
	if arguments.Year < 2015 || arguments.Year > time.Now().Year() {
		fmt.Printf("Year out of range 2015 -  %d input:%d\n", time.Now().Year(), arguments.Year)
		os.Exit(1)
	}

	// Validate the day input
	if arguments.Day < 1 || arguments.Day > 25 {
		fmt.Printf("Day must be a number between 1-25")
		os.Exit(1)
	}

	currWd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error while getting current working dir: %v\n", err)
		os.Exit(1)
	}
	path := fmt.Sprintf("%s/solutions/Y_%d/D_%d", currWd, arguments.Year, arguments.Day)
	// First check if we have the folder for the year we are downloading in the format Y_XXXX
	if err := os.MkdirAll(path, 0777); err != nil {
		fmt.Printf("Error creating dir: %v\n", err)
		os.Exit(1)
	}

	client := grab.NewClient()

	// start download
	req, err := grab.NewRequest(fmt.Sprintf("%s/day_%d_input.txt", path, arguments.Day), fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", arguments.Year, arguments.Day))
	fmt.Printf("Downloading %v ...\n", req.URL())
	req.HTTPRequest.Header.Set("Cookie", fmt.Sprintf("session=%s", os.Getenv("SESSION_KEY")))
	resp := client.Do(req)
	//fmt.Printf("  %v\n", resp.HTTPResponse.Status)
	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size(),
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Printf("Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
	fmt.Println("Creating go module and test from template")
	moduleContent, err := os.ReadFile("./templates/day_x.go")
	moduleContent = bytes.ReplaceAll(moduleContent, []byte("YX"), []byte(fmt.Sprintf("%d", arguments.Year)))
	moduleContent = bytes.ReplaceAll(moduleContent, []byte("DX"), []byte(fmt.Sprintf("%d", arguments.Day)))
	if err := os.WriteFile(fmt.Sprintf("%s/day_%d.go", path, arguments.Day), moduleContent, 0777); err != nil {
		fmt.Printf("Creating files from template failed: %v\n", err)
		os.Exit(1)
	}

	testContent, err := os.ReadFile("./templates/day_x_test.go")
	testContent = bytes.ReplaceAll(testContent, []byte("YX"), []byte(fmt.Sprintf("%d", arguments.Year)))
	testContent = bytes.ReplaceAll(testContent, []byte("DX"), []byte(fmt.Sprintf("%d", arguments.Day)))
	if err := os.WriteFile(fmt.Sprintf("%s/day_%d_test.go", path, arguments.Day), testContent, 0777); err != nil {
		fmt.Printf("Creating files from template failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Templates created have fun with todays puzzle!")
}

func ReadFullFileInput(fileStruct FileStruct) ([]byte, error) {
	// Either we get a path (actual input) or an in-memory []byte
	if len(fileStruct.Contents) > 0 {
		return fileStruct.Contents, nil
	}
	return os.ReadFile(fileStruct.Path)
}

func ParseLinesFromFullInput(fileContents []byte) []string {
	// Either we get a path (actual input) or an in-memory []byte
	return regexp.MustCompile(`\r?\n`).Split(string(fileContents), -1)
}

func MustParseStringToInt(integer string) int {
	i, err := strconv.Atoi(integer)
	if err != nil {
		fmt.Printf("The following string could not be parsed to an int: %s", integer)
		os.Exit(1)
	}
	return i
}

func ParseStringSliceToUInt64Slice(slice []string) (output []uint64) {
	for _, item := range slice {
		output = append(output, uint64(MustParseStringToInt(item)))
	}
	return output
}

func RuneIntersectionUnique(s1 []rune, s2 []rune) []rune {
	var intersection []rune
	charMap := make(map[rune]bool, len(s1))
	// Setup a map
	for _, item := range s1 {
		charMap[item] = true
	}
	// IF we have the same char append as intersection
	for _, item := range s2 {
		if charMap[item] {
			// Dedup on the GO
			if !RuneSliceContains(intersection, item) {
				intersection = append(intersection, item)
			}
		}
	}
	return intersection
}

func RuneSliceContains(slice []rune, item rune) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func ReverseStringSlice(s []string) []string {
	a := make([]string, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func ReverseIntSlice(s []int) []int {
	a := make([]int, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func PopFirstItemUInt64Slice(s []uint64) (item uint64, slice []uint64) {
	if len(s) == 0 {
		fmt.Print("Cannot pop item from an empty slice")
		os.Exit(1)
	} else if len(s) == 1 {
		return s[0], s[0:0]
	}
	return s[0], s[1:]
}

func ReverseMatrixSlices(matrix [][]int) [][]int {
	var rowMatrix [][]int
	for _, row := range matrix {
		rowMatrix = append(rowMatrix, ReverseIntSlice(row))
	}

	return rowMatrix
}

func ReverseMatrixColumnsAndRows(matrix [][]int) [][]int {
	var colMatrix [][]int
	for rIndex, row := range matrix {
		for cIndex, col := range row {
			if rIndex == 0 { // only append the first row as new column rows
				colMatrix = append(colMatrix, []int{})
			}
			colMatrix[cIndex] = append(colMatrix[cIndex], col)
		}
	}
	return colMatrix
}
