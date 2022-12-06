package main

import (
	"fmt"
	"testing"
)

func getEmptyCrates() [][]string {
	return [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}}
}

func TestFindMarker(t *testing.T) {
	cases := []struct {
		input    string
		token    string
		expected string
	}{
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			token:    "",
			expected: "vwbj",
		},
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			token:    "",
			expected: "jpqm",
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			token:    "",
			expected: "pdvj",
		},
	}

	for _, c := range cases {
		tmpMarker := c.token
		ri := 0
		for _, v := range c.input {
			unique, marker := FindMarker(string(v), tmpMarker)
			tmpMarker = marker
			ri++
			if unique {
				break
			}
		}

		if tmpMarker != c.expected {
			t.Error(fmt.Sprintf("Fail. Expected: %v but got: %v", c.expected, tmpMarker))
		}
	}
}
