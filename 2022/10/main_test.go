package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	cases := []struct {
		input    string
		expected struct {
			op  string
			val int
		}
	}{
		{
			input: "noop",
			expected: struct {
				op  string
				val int
			}{op: "noop", val: 0},
		},
		{
			input: "addx -3",
			expected: struct {
				op  string
				val int
			}{op: "addx", val: -3},
		},
		{
			input: "addx 5",
			expected: struct {
				op  string
				val int
			}{op: "addx", val: 5},
		},
	}

	for _, c := range cases {
		op, val := Read(c.input)
		if op != c.expected.op && val != c.expected.val {
			errStr := fmt.Sprintf("Fail. Expected: %v and %v but got: %v and %v", c.expected.op, c.expected.val, op,
				val)
			t.Error(errStr)
		}
	}
}

func TestExecute(t *testing.T) {
	cases := []struct {
		input     []string
		execution *Execution
		expected  *Execution
	}{
		{
			input: testInput1,
			execution: &Execution{
				registerX:       1,
				operation:       Operation{},
				cycle:           0,
				multiplierIndex: 0,
				strength:        0,
			},
			expected: &Execution{
				registerX:       20,
				operation:       Operation{},
				cycle:           20,
				multiplierIndex: 1,
				strength:        420,
			},
		},
		{
			input: testInput2,
			execution: &Execution{
				registerX:       1,
				operation:       Operation{},
				cycle:           0,
				multiplierIndex: 0,
				strength:        0,
			},
			expected: &Execution{
				registerX:       1,
				operation:       Operation{},
				cycle:           20,
				multiplierIndex: 1,
				strength:        13140,
			},
		},
	}

	for _, c := range cases {
		for _, v := range c.input {
			op, val := Read(v)
			c.execution.operation = Operation{
				name: op,
				val:  val,
			}
			c.execution.line = v
			Execute(c.execution)
		}

		if c.execution.strength != c.expected.strength {
			errStr := fmt.Sprintf("Fail.\nExp -> Strength: %v\nGot -> Strength: %v", c.expected.strength, c.execution.strength)
			t.Error(errStr)
		}
	}
}

func TestDisplayOnCrt(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
	}{
		{
			input: testInput2,
			expected: []string{
				"##..##..##..##..##..##..##..##..##..##..",
				"###...###...###...###...###...###...###.",
				"####....####....####....####....####....",
				"#####.....#####.....#####.....#####.....",
				"######......######......######......####",
				"#######.......#######.......#######.....",
			},
		},
	}

	for _, c := range cases {
		crt := &CRT{
			spritePos: 1,
			rows:      []string{"", "", "", "", "", ""},
		}
		for _, v := range c.input {
			op, val := Read(v)
			DisplayOnCrt(crt, Operation{
				name: op,
				val:  val,
			})
		}

		if !reflect.DeepEqual(c.expected, crt.rows) {
			fmt.Println("Expected:")
			for _, v := range c.expected {
				fmt.Println(v)
			}
			fmt.Println("Result:")
			for _, v := range crt.rows {
				fmt.Println(v)
			}
			t.Error("Failed!")
		}
	}
}

var (
	testInput1 = []string{
		"addx 15",
		"addx -11",
		"addx 6",
		"addx -3",
		"addx 5",
		"addx -1",
		"addx -8",
		"addx 13",
		"addx 4",
		"noop",
		"addx -1",
	}

	testInput2 = []string{
		"addx 15",
		"addx -11",
		"addx 6",
		"addx -3",
		"addx 5",
		"addx -1",
		"addx -8",
		"addx 13",
		"addx 4",
		"noop",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx 5",
		"addx -1",
		"addx -35",
		"addx 1",
		"addx 24",
		"addx -19",
		"addx 1",
		"addx 16",
		"addx -11",
		"noop",
		"noop",
		"addx 21",
		"addx -15",
		"noop",
		"noop",
		"addx -3",
		"addx 9",
		"addx 1",
		"addx -3",
		"addx 8",
		"addx 1",
		"addx 5",
		"noop",
		"noop",
		"noop",
		"noop",
		"noop",
		"addx -36",
		"noop",
		"addx 1",
		"addx 7",
		"noop",
		"noop",
		"noop",
		"addx 2",
		"addx 6",
		"noop",
		"noop",
		"noop",
		"noop",
		"noop",
		"addx 1",
		"noop",
		"noop",
		"addx 7",
		"addx 1",
		"noop",
		"addx -13",
		"addx 13",
		"addx 7",
		"noop",
		"addx 1",
		"addx -33",
		"noop",
		"noop",
		"noop",
		"addx 2",
		"noop",
		"noop",
		"noop",
		"addx 8",
		"noop",
		"addx -1",
		"addx 2",
		"addx 1",
		"noop",
		"addx 17",
		"addx -9",
		"addx 1",
		"addx 1",
		"addx -3",
		"addx 11",
		"noop",
		"noop",
		"addx 1",
		"noop",
		"addx 1",
		"noop",
		"noop",
		"addx -13",
		"addx -19",
		"addx 1",
		"addx 3",
		"addx 26",
		"addx -30",
		"addx 12",
		"addx -1",
		"addx 3",
		"addx 1",
		"noop",
		"noop",
		"noop",
		"addx -9",
		"addx 18",
		"addx 1",
		"addx 2",
		"noop",
		"noop",
		"addx 9",
		"noop",
		"noop",
		"noop",
		"addx -1",
		"addx 2",
		"addx -37",
		"addx 1",
		"addx 3",
		"noop",
		"addx 15",
		"addx -21",
		"addx 22",
		"addx -6",
		"addx 1",
		"noop",
		"addx 2",
		"addx 1",
		"noop",
		"addx -10",
		"noop",
		"noop",
		"addx 20",
		"addx 1",
		"addx 2",
		"addx 2",
		"addx -6",
		"addx -11",
		"noop",
		"noop",
		"noop",
	}
)
