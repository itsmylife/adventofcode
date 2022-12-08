package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindScenicScore(t *testing.T) {
	cases := []struct {
		input    []string
		expected float64
	}{
		{
			input: []string{
				"30373",
				"25512",
				"65332",
				"33549",
				"35390",
			},
			expected: 8,
		},
	}

	for _, c := range cases {
		result := FindScenicScore(c.input)
		if result != c.expected {
			t.Error(fmt.Sprintf("Fail. Expected: %v but got: %v", c.expected, result))
		}
	}
}

func TestFindVisibleTreeCount(t *testing.T) {
	cases := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"30373",
				"25512",
				"65332",
				"33549",
				"35390",
			},
			expected: 21,
		},
	}

	for _, c := range cases {
		result := FindVisibleTreeCount(c.input)
		if result != c.expected {
			t.Error(fmt.Sprintf("Fail. Expected: %v but got: %v", c.expected, result))
		}
	}
}

func TestCalculateHighest(t *testing.T) {
	cases := []struct {
		input    []string
		expected struct {
			rows []int64
			cols []int64
		}
	}{
		{
			input: []string{
				"30373",
				"25512",
				"65332",
				"33549",
				"35390",
			},
			expected: struct {
				rows []int64
				cols []int64
			}{rows: []int64{7, 5, 6, 9, 9}, cols: []int64{6, 5, 5, 9, 9}},
		},
	}

	for _, c := range cases {
		var rows = make([]int64, 5)
		var cols = make([]int64, 5)
		for i, v := range c.input {
			rows, cols = CalculateHighest(v, i, rows, cols)
		}
		if !reflect.DeepEqual(rows, c.expected.rows) {
			t.Error(fmt.Sprintf("Rows failed. Expected: %v but got: %v", c.expected.rows, rows))
		}
		if !reflect.DeepEqual(cols, c.expected.cols) {
			t.Error(fmt.Sprintf("Cols failed. Expected: %v but got: %v", c.expected.cols, cols))
		}
	}
}
