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
			args: args{inputStruct: utils.FileStruct{Contents: []byte("$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k")}},
			want: 95437,
		},
		{
			name: "Part 1 live input",
			args: args{inputStruct: utils.FileStruct{Path: "day_7_input.txt"}},
			want: 1477771,
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
			args: args{inputStruct: utils.FileStruct{Contents: []byte("$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k")}},
			want: 24933642,
		},
		{
			name: "Part 2 live input",
			args: args{inputStruct: utils.FileStruct{Path: "day_7_input.txt"}},
			want: 3579501,
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
