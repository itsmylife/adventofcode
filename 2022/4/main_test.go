package main

import (
	"fmt"
	"testing"
)

func TestFindContainedRanges(t *testing.T) {
	cases := []struct {
		line     string
		expected int
	}{
		{line: "2-8,3-7", expected: 1},
		{line: "5-7,7-9", expected: 0},
		{line: "6-6,4-6", expected: 1},
		{line: "2-6,4-8", expected: 0},
	}

	for _, c := range cases {
		point := FindContainedRanges(c.line)

		if point != c.expected {
			t.Error(fmt.Sprintf("Line: %s failed. Expected: %d but got: %d", c.line, c.expected, point))
		}
	}
}

func TestFindOverlappingRangespingRanges(t *testing.T) {
	cases := []struct {
		line     string
		expected int
	}{
		{line: "2-8,3-7", expected: 1},
		{line: "5-7,7-9", expected: 1},
		{line: "6-6,4-6", expected: 1},
		{line: "2-6,4-8", expected: 1},
		{line: "2-6,7-8", expected: 0},
		{line: "1-1,2-2", expected: 0},
		{line: "1-5,8-9", expected: 0},
		{line: "68-69,45-69", expected: 1},
		{line: "68-69,45-67", expected: 0},
	}

	for _, c := range cases {
		point := FindOverlappingRanges(c.line)

		if point != c.expected {
			t.Error(fmt.Sprintf("Line: %s failed. Expected: %d but got: %d", c.line, c.expected, point))
		}
	}
}
