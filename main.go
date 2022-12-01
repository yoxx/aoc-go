package main

import (
	"aoc-go/utils"
	"fmt"
	"os"
)

func main() {
	arguments := utils.ParseCLIArguments()
	switch arguments.Command {
	case "download":
		utils.DownloadInput(arguments)
		break
	case "puzzle":
		// Run the puzzle code
		utils.RunPuzzle(arguments)
		break
	case "tests":
		utils.TestPuzzle(arguments)
		break
	case "default":
		fmt.Printf("You must provide a -command, see -help for options.")
		os.Exit(1)
	}
}
