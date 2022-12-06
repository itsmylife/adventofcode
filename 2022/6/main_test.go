package main

import (
	"fmt"
	"testing"
)

func getEmptyCrates() [][]string {
	return [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}}
}

func TestFindMarkerAfter4(t *testing.T) {
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
			unique, marker := FindMarker(string(v), tmpMarker, 4)
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

func TestFindMarkerAfter14(t *testing.T) {
	cases := []struct {
		input    string
		token    string
		expected string
	}{
		{
			input:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			token:    "",
			expected: "qmgbljsphdztnv",
		},
		{
			input:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			token:    "",
			expected: "vbhsrlpgdmjqwf",
		},
		{
			input:    "nppdvjthqldpwncqszvftbrmjlhg",
			token:    "",
			expected: "ldpwncqszvftbr",
		},
	}

	for _, c := range cases {
		tmpMarker := c.token
		ri := 0
		for _, v := range c.input {
			unique, marker := FindMarker(string(v), tmpMarker, 14)
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
