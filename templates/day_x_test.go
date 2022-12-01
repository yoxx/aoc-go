package main

import (
	"fmt"
	"testing"
)

func TestPartOne(t *testing.T) {
	fmt.Println("Testing Part One")
	expected := 0
	output := PartOne()
	// TODO write a test bro
	if output != expected {
		t.Errorf("expected %d for the test file got:", expected)
	}
}

func TestPartTwo(t *testing.T) {
	fmt.Println("Testing Part Two")
	expected := 0
	output := PartTwo()
	// TODO write a test bro
	if output != expected {
		t.Errorf("expected %d for the test file got:", expected)
	}
}
