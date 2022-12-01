package main

import (
	"aoc-go/utils"
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	fmt.Println("YX_DX Testing Part One")
	expected := 0
	inputStruct := utils.FileStruct{} // TODO contents!
	output := PartOne(inputStruct)
	// TODO write a test bro
	if output != expected {
		t.Errorf("expected %d for the test file got:", expected)
	}
}

func TestPartTwo(t *testing.T) {
	fmt.Println("YX_DX Testing Part Two")
	expected := 0
	inputStruct := utils.FileStruct{} // TODO contents!
	output := PartTwo(inputStruct)
	// TODO write a test bro
	if output != expected {
		t.Errorf("expected %d for the test file got:", expected)
	}
}
