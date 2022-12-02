package main

import (
	"aoc-go/utils"
	"testing"
)

func TestPartOne(t *testing.T) {
	type args struct {
		inputStruct utils.FileStruct
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Part 1 test",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("A Y\nB X\nC Z")}},
			want: 15,
		},
		{
			name: "Part 1 live input",
			args: args{inputStruct: utils.FileStruct{Path: "day_2_input.txt"}},
			want: 13675,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartOne(tt.args.inputStruct); got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	type args struct {
		inputStruct utils.FileStruct
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Part 2 test",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("A Y\nB X\nC Z")}},
			want: 12,
		},
		{
			name: "Part 2 live input",
			args: args{inputStruct: utils.FileStruct{Path: "day_2_input.txt"}},
			want: 14184,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.inputStruct); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
