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
			args: args{inputStruct: utils.FileStruct{Contents: []byte("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet")}},
			want: 142,
		},
		{
			name: "Part 1 test double digits",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("11abc2\npqr3stu88vwx\na1b2c3d4e5f\ntreb7uchet")}},
			want: 142,
		},
		{
			name: "Part 1 live input",
			args: args{inputStruct: utils.FileStruct{Path: "day_1_input.txt"}},
			want: 54605,
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
			args: args{inputStruct: utils.FileStruct{Contents: []byte("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")}},
			want: 281,
		},
		{
			name: "Part 2 test eighthree",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("eighthree")}},
			want: 83,
		},
		{
			name: "Part 2 test sevenine",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("sevenine")}},
			want: 79,
		},
		//{
		//	name: "Part 2 live input",
		//	args: args{inputStruct: utils.FileStruct{Path: "day_1_input.txt"}},
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
