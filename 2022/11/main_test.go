package main

import (
	"fmt"
	"reflect"
	"testing"
)

// func TestCalculateTheWorryLevel(t *testing.T) {
// 	cases := []struct {
// 		input              *Monkey
// 		expectedWorry      int
// 		expectedNextMonkey int
// 	}{
// 		{
// 			input: &Monkey{
// 				items: []int{10},
// 				throwMap: map[bool]int{
// 					true:  3,
// 					false: 5,
// 				},
// 				test: 5,
// 				operation: &Operation{
// 					input1: "old",
// 					op:     "+",
// 					input2: "3",
// 				},
// 			},
// 			expectedWorry:      4,
// 			expectedNextMonkey: 5,
// 		},
// 	}
//
// 	for _, c := range cases {
// 		result := CalculateTheWorryLevel(c.input, c.input.items[0], true)
// 		if result != c.expectedWorry {
// 			errStr := fmt.Sprintf("Fail. Expected: %v but got: %v", c.expectedWorry, result)
// 			t.Error(errStr)
// 		}
// 	}
// }

func TestRunTheSimulation(t *testing.T) {
	cases := []struct {
		input    []string
		rounds   int
		relief   bool
		expected []int
	}{
		{
			input:    testInput,
			rounds:   20,
			relief:   false,
			expected: []int{99, 97, 8, 103},
		},
		{
			input:    testInput,
			rounds:   20,
			relief:   true,
			expected: []int{101, 95, 7, 105},
		},
		{
			input:    testInput,
			rounds:   10000,
			relief:   false,
			expected: []int{52166, 47830, 1938, 52013},
		},
	}

	for _, c := range cases {
		var monkeys []*Monkey
		for i, line := range c.input {
			monkeys = ParseInput(line, i+1, monkeys)
		}
		biznis := RunTheSimulation(monkeys, c.rounds, c.relief)
		if !reflect.DeepEqual(c.expected, biznis) {
			errStr := fmt.Sprintf("Fail. \nExpected: \n%v \nbut got: \n%v", c.expected, biznis)
			t.Error(errStr)
		}
	}
}

var testInput = []string{
	"Monkey 0:",
	"Starting items: 79, 98",
	"Operation: new = old * 19",
	"Test: divisible by 23",
	"If true: throw to monkey 2",
	"If false: throw to monkey 3",
	"",
	"Monkey 1:",
	"Starting items: 54, 65, 75, 74",
	"Operation: new = old + 6",
	"Test: divisible by 19",
	"If true: throw to monkey 2",
	"If false: throw to monkey 0",
	"",
	"Monkey 2:",
	"Starting items: 79, 60, 97",
	"Operation: new = old * old",
	"Test: divisible by 13",
	"If true: throw to monkey 1",
	"If false: throw to monkey 3",
	"",
	"Monkey 3:",
	"Starting items: 74",
	"Operation: new = old + 3",
	"Test: divisible by 17",
	"If true: throw to monkey 0",
	"If false: throw to monkey 1",
}
