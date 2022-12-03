package main

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	cases := []struct{ index1, index2, expected int }{
		{index1: 0, index2: 0, expected: 4}, // rock 		rock		draw
		{index1: 1, index2: 1, expected: 5}, // paper 		paper		draw
		{index1: 2, index2: 2, expected: 6}, // scissors 	scissors	draw
		{index1: 2, index2: 0, expected: 7}, // scissors 	rock		win
		{index1: 0, index2: 1, expected: 8}, // rock 		paper 		win
		{index1: 1, index2: 2, expected: 9}, // paper 		scissors	win
		{index1: 1, index2: 0, expected: 1}, // paper 		rock		lost
		{index1: 2, index2: 1, expected: 2}, // scissors 	paper		lost
		{index1: 0, index2: 2, expected: 3}, // rock 		scissors	lost
	}

	for _, c := range cases {
		point := CalculateRockPaperScissors(c.index1, c.index2)

		if point != c.expected {
			t.Error(fmt.Sprintf("Index1: %d -- Index2: %d failed. Expected: %d but got: %d", c.index1, c.index2, c.expected, point))
		}
	}
}

func TestCalculateWhatsToPlay(t *testing.T) {
	cases := []struct {
		cols     []string
		expected int
	}{
		{cols: []string{"A", "X"}, expected: 3}, // rock lost (scissors 3) 3
		{cols: []string{"A", "Y"}, expected: 4}, // rock draw (rock 1) 4
		{cols: []string{"A", "Z"}, expected: 8}, // rock win (paper 2) 8
		{cols: []string{"B", "X"}, expected: 1}, // paper lost (rock 1) 1
		{cols: []string{"B", "Y"}, expected: 5}, // paper draw (paper 2) 5
		{cols: []string{"B", "Z"}, expected: 9}, // paper win (scissors 3) 9
		{cols: []string{"C", "X"}, expected: 2}, // scissors lost (paper 2) 2
		{cols: []string{"C", "Y"}, expected: 6}, // scissors draw (scissors 3) 6
		{cols: []string{"C", "Z"}, expected: 7}, // scissors win (rock 1) 7

	}

	for _, c := range cases {
		point := CalculateWhatsToPlay(c.cols)

		if point != c.expected {
			t.Error(fmt.Sprintf("Index1: %s -- Index2: %s failed. Expected: %d but got: %d", c.cols[0], c.cols[1], c.expected, point))
		}
	}
}
