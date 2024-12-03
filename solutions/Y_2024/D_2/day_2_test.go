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
			args: args{inputStruct: utils.FileStruct{Contents: []byte("7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9")}},
			want: 2,
		},
		{
			name: "Part 1 live input",
			args: args{inputStruct: utils.FileStruct{Path: "day_2_input.txt"}},
			want: 334,
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
			args: args{inputStruct: utils.FileStruct{Contents: []byte("7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9")}},
			want: 4,
		},
		{
			name: "Part 2 test - first number wrong to high difference",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("9 3 4 5 6")}},
			want: 1,
		},
		{
			name: "Part 2 test - first number wrong to wrong direction",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("4 3 4 5 6")}},
			want: 1,
		},
		{
			name: "Part 2 test - last number wrong to wrong direction",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("2 3 4 5 4")}},
			want: 1,
		},
		{
			name: "Part 2 test - last number wrong to high difference",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("2 3 4 5 1")}},
			want: 1,
		},
		//{
		//	name: "Part 2 live input",
		//	args: args{inputStruct: utils.FileStruct{Path: "day_2_input.txt"}},
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
