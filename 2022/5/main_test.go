package main

import (
	"fmt"
	"reflect"
	"testing"
)

func getEmptyCrates() [][]string {
	return [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}}
}

func TestReadInitialCrates(t *testing.T) {
	cases := []struct {
		line     string
		expected [][]string
	}{
		{
			line: "        [H]         [S]         [D]",
			expected: [][]string{{}, {}, {"[H]"}, {}, {}, {"[S]"}, {}, {},
				{"[D]"}},
		},
		{
			line: "    [S] [C]         [C]     [Q] [L]",
			expected: [][]string{{}, {"[S]"}, {"[C]"}, {}, {}, {"[C]"}, {}, {"[Q]"},
				{"[L]"}},
		},
		{
			line: "[H] [Q] [P] [L] [G] [V] [Z] [D] [B]",
			expected: [][]string{{"[H]"}, {"[Q]"}, {"[P]"}, {"[L]"}, {"[G]"}, {"[V]"}, {"[Z]"}, {"[D]"},
				{"[B]"}},
		},
	}

	for _, c := range cases {
		initialCrates := getEmptyCrates()
		ReadInitialCrates(c.line, initialCrates)

		if !reflect.DeepEqual(initialCrates, c.expected) {
			t.Error(fmt.Sprintf("Line: %s failed. Expected: %v but got: %v", c.line, c.expected, initialCrates))
		}
	}
}

func TestMoveCratesWithCrane9000(t *testing.T) {
	cases := []struct {
		line     string
		crates   [][]string
		expected [][]string
	}{
		{
			line:   "move 2 from 7 to 2",
			crates: [][]string{{"[H]"}, {"[Q]"}, {"[P]"}, {"[L]"}, {"[G]"}, {"[V]"}, {"[Z]", "[C]"}, {"[D]"}, {"[B]"}},
			expected: [][]string{{"[H]"}, {"[Q]", "[C]", "[Z]"}, {"[P]"}, {"[L]"}, {"[G]"}, {"[V]"}, {}, {"[D]"},
				{"[B]"}},
		},
	}

	for _, c := range cases {

		MoveCratesWithCrane9000(c.line, c.crates)

		if !reflect.DeepEqual(c.crates, c.expected) {
			t.Error(fmt.Sprintf("Line: %s failed. Expected: %v but got: %v", c.line, c.expected, c.crates))
		}
	}
}

func TestMoveCratesWithCrane9001(t *testing.T) {
	cases := []struct {
		line     string
		crates   [][]string
		expected [][]string
	}{
		{
			line:   "move 2 from 7 to 2",
			crates: [][]string{{"[H]"}, {"[Q]"}, {"[P]"}, {"[L]"}, {"[G]"}, {"[V]"}, {"[Z]", "[C]"}, {"[D]"}, {"[B]"}},
			expected: [][]string{{"[H]"}, {"[Q]", "[Z]", "[C]"}, {"[P]"}, {"[L]"}, {"[G]"}, {"[V]"}, {}, {"[D]"},
				{"[B]"}},
		},
	}

	for _, c := range cases {

		MoveCratesWithCrane9001(c.line, c.crates)

		if !reflect.DeepEqual(c.crates, c.expected) {
			t.Error(fmt.Sprintf("Line: %s failed. Expected: %v but got: %v", c.line, c.expected, c.crates))
		}
	}
}

func TestReadTopCrates(t *testing.T) {
	cases := []struct {
		crates   [][]string
		expected string
	}{
		{
			crates: [][]string{{"[H]"}, {"[Q]"}, {"[P]"}, {"[L]"}, {"[G]"}, {"[V]"}, {"[Z]"}, {"[D]"},
				{"[B]"}},
			expected: "HQPLGVZDB",
		},
	}

	for _, c := range cases {
		word := ReadTopCrates(c.crates)

		if word != c.expected {
			t.Error(fmt.Sprintf("Crates: %v failed. Expected: %v but got: %v", c.crates, c.expected, word))
		}
	}
}
