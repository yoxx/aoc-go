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
			args: args{inputStruct: utils.FileStruct{Contents: []byte("Monkey 0:\n  Starting items: 79, 98\n  Operation: new = old * 19\n  Test: divisible by 23\n    If true: throw to monkey 2\n    If false: throw to monkey 3\n\nMonkey 1:\n  Starting items: 54, 65, 75, 74\n  Operation: new = old + 6\n  Test: divisible by 19\n    If true: throw to monkey 2\n    If false: throw to monkey 0\n\nMonkey 2:\n  Starting items: 79, 60, 97\n  Operation: new = old * old\n  Test: divisible by 13\n    If true: throw to monkey 1\n    If false: throw to monkey 3\n\nMonkey 3:\n  Starting items: 74\n  Operation: new = old + 3\n  Test: divisible by 17\n    If true: throw to monkey 0\n    If false: throw to monkey 1")}},
			want: 10605,
		},
		{
			name: "Part 1 live input",
			args: args{inputStruct: utils.FileStruct{Path: "day_11_input.txt"}},
			want: 119715,
		},
		{
			name: "Part 1 live input dick",
			args: args{inputStruct: utils.FileStruct{Contents: []byte("Monkey 0:\n  Starting items: 91, 58, 52, 69, 95, 54\n  Operation: new = old * 13\n  Test: divisible by 7\n    If true: throw to monkey 1\n    If false: throw to monkey 5\n\nMonkey 1:\n  Starting items: 80, 80, 97, 84\n  Operation: new = old * old\n  Test: divisible by 3\n    If true: throw to monkey 3\n    If false: throw to monkey 5\n\nMonkey 2:\n  Starting items: 86, 92, 71\n  Operation: new = old + 7\n  Test: divisible by 2\n    If true: throw to monkey 0\n    If false: throw to monkey 4\n\nMonkey 3:\n  Starting items: 96, 90, 99, 76, 79, 85, 98, 61\n  Operation: new = old + 4\n  Test: divisible by 11\n    If true: throw to monkey 7\n    If false: throw to monkey 6\n\nMonkey 4:\n  Starting items: 60, 83, 68, 64, 73\n  Operation: new = old * 19\n  Test: divisible by 17\n    If true: throw to monkey 1\n    If false: throw to monkey 0\n\nMonkey 5:\n  Starting items: 96, 52, 52, 94, 76, 51, 57\n  Operation: new = old + 3\n  Test: divisible by 5\n    If true: throw to monkey 7\n    If false: throw to monkey 3\n\nMonkey 6:\n  Starting items: 75\n  Operation: new = old + 5\n  Test: divisible by 13\n    If true: throw to monkey 4\n    If false: throw to monkey 2\n\nMonkey 7:\n  Starting items: 83, 75\n  Operation: new = old + 1\n  Test: divisible by 19\n    If true: throw to monkey 2\n    If false: throw to monkey 6\n")}},
			want: 57348,
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
			args: args{inputStruct: utils.FileStruct{Contents: []byte("Monkey 0:\n  Starting items: 79, 98\n  Operation: new = old * 19\n  Test: divisible by 23\n    If true: throw to monkey 2\n    If false: throw to monkey 3\n\nMonkey 1:\n  Starting items: 54, 65, 75, 74\n  Operation: new = old + 6\n  Test: divisible by 19\n    If true: throw to monkey 2\n    If false: throw to monkey 0\n\nMonkey 2:\n  Starting items: 79, 60, 97\n  Operation: new = old * old\n  Test: divisible by 13\n    If true: throw to monkey 1\n    If false: throw to monkey 3\n\nMonkey 3:\n  Starting items: 74\n  Operation: new = old + 3\n  Test: divisible by 17\n    If true: throw to monkey 0\n    If false: throw to monkey 1")}},
			want: 2713310158,
		},
		//{
		//	name: "Part 2 live input",
		//	args: args{inputStruct: utils.FileStruct{Path: "day_11_input.txt"}},
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
