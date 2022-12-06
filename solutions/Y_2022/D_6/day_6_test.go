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
			name: "Part 1 test v1",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb")}},
			want: 7,
		},
		{
			name: "Part 1 test v2",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("bvwbjplbgvbhsrlpgdmjqwftvncz")}},
			want: 5,
		},
		{
			name: "Part 1 test v3",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("nppdvjthqldpwncqszvftbrmjlhg")}},
			want: 6,
		},
		{
			name: "Part 1 test v4",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")}},
			want: 10,
		},
		{
			name: "Part 1 test v5",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")}},
			want: 11,
		},
		{
			name: "Part 1 live input",
			args: args{inputStruct: utils.FileStruct{Path: "day_6_input.txt"}},
			want: 1640,
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
			name: "Part 1 test v1",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb")}},
			want: 19,
		},
		{
			name: "Part 1 test v2",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("bvwbjplbgvbhsrlpgdmjqwftvncz")}},
			want: 23,
		},
		{
			name: "Part 1 test v3",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("nppdvjthqldpwncqszvftbrmjlhg")}},
			want: 23,
		},
		{
			name: "Part 1 test v4",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")}},
			want: 29,
		},
		{
			name: "Part 1 test v5",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")}},
			want: 26,
		},
		{
			name: "Part 2 live input",
			args: args{inputStruct: utils.FileStruct{Path: "day_6_input.txt"}},
			want: 3613,
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
