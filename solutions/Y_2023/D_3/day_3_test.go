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
			args: args{inputStruct: utils.FileStruct{Contents: []byte("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..")}},
			want: 4361,
		},
		{
			name: "Part 1 live input",
			args: args{inputStruct: utils.FileStruct{Path: "day_3_input.txt"}},
			want: 0,
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
			args: args{inputStruct: utils.FileStruct{Contents: []byte("TODO_FILL_YOUR_CONTENTS_HERE")}},
			want: 0,
		},
		//{
		//	name: "Part 2 live input",
		//	args: args{inputStruct: utils.FileStruct{Path: "day_3_input.txt"}},
		//	want: 0,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.inputStruct); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
