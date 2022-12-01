# Advent-Of-Code in Go
Advent of Code is an Advent calendar of small programming puzzles for a variety of
skill sets and skill levels that can be solved in any programming language you like.

For more information about the advent of code see [adventofcode.com](https://adventofcode.com)

## General Info
* This project is written in Golang
    * Current Go version 1.19.3
* I like to do the advent puzzles for as long as I have time, usually it gets busier during the december month.
  This will result in me picking my battles and focussing on other things. I do it for fun anyways.
* Previous years I have tried to do my AOC in PHP. The repo for that can be found here [advent-of-code](https://github.com/yoxx/advent-of-code)

## Pipelines Wishlist
* Should have a github actions pipeline that does the following:
  * Runs a formatter when executing the pipeline
  * Check for security updates and such
  * Runs the unittests

## Startup
 * To get your repo working copy the .env.example to a .env file
 * Fill the session key with the session key in your AOC cookie (on the site)
 * Download input for a certain day with the following command
   * This code will download the AOC input of that day and create a template go file as wel as a test file
   * These files can be edited to include the code required to make the solution work
      ```go run main.go -command download -year 2022 -day 1```
 * Run the generated AOC day for a certain day with the following command
   * Optionally include the `-part` flag to run either part 1 or 2
     ```go run main.go -command puzzle -year 2022 -day 1```
 * Run the unittests for a certain day with the following command
   * Optionally include the `-part` flag to run either part 1 or 2
   ```go run main.go -command tests -year 2022 -day 1```

##Years
### 2022
* First started with the Advent of Code in Go
* Setup some utils for running puzzles and downloading input.
* Wishes: Function to get the leaderboard