package main

import (
	"fmt"
	"testing"
)

func TestCalculatePriorityPoint(t *testing.T) {
	cases := []struct {
		line     string
		expected int
	}{
		{line: "AA", expected: 27},                       // A
		{line: "asedtGDEen", expected: 5},                // e
		{line: "vJrwpWtwJgWrhcsFMMfFFhFp", expected: 16}, // p
		{line: "ttgJtRGJQctTZtZT", expected: 20},         // t
		{line: "PmmdzqPrVvPwwTWBwg", expected: 42},       // P
	}

	for _, c := range cases {
		point := CalculatePriorityPoint(c.line)

		if point != c.expected {
			t.Error(fmt.Sprintf("Line: %s failed. Expected: %d but got: %d", c.line, c.expected, point))
		}
	}
}

func TestCalculateGroupPriorityPoint(t *testing.T) {
	cases := []struct {
		lines    []string
		expected int
	}{
		{lines: []string{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
		}, expected: 18}, // r
		{lines: []string{
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		}, expected: 52}, // Z
	}

	for _, c := range cases {
		point := CalculateGroupPriorityPoint(c.lines)

		if point != c.expected {
			t.Error(fmt.Sprintf("Line: %s failed. Expected: %d but got: %d", c.lines, c.expected, point))
		}
	}
}
