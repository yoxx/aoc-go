package main

import (
	"aoc-go/utils"
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	fmt.Println("2022_1 Testing Part One")
	expected := 24000
	inputStruct := utils.FileStruct{
		Contents: []byte("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"),
	}
	output := PartOne(inputStruct)
	// TODO write a test bro
	if output != expected {
		t.Errorf("expected %d for the test got:", expected)
	}
}

func TestPartTwo(t *testing.T) {
	fmt.Println("2022_1 Testing Part Two")
	expected := 45000
	inputStruct := utils.FileStruct{
		Contents: []byte("1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"),
	}
	output := PartTwo(inputStruct)
	// TODO write a test bro
	if output != expected {
		t.Errorf("expected %d for the test file got:", expected)
	}
}
