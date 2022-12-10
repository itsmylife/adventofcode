package main

import (
	"fmt"
	"testing"
)

func TestShouldGoDiagonal(t *testing.T) {
	cases := []struct {
		input struct {
			hp, tp *Pos
		}
		expected bool
	}{
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 0,
				y: 0,
			}, tp: &Pos{
				x: 0,
				y: 0,
			}},
			expected: false,
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 1,
				y: 0,
			}, tp: &Pos{
				x: 0,
				y: 0,
			}},
			expected: false,
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 1,
				y: 5,
			}, tp: &Pos{
				x: -1,
				y: 0,
			}},
			expected: true,
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 4,
				y: 4,
			}, tp: &Pos{
				x: 0,
				y: 0,
			}},
			expected: true,
		},
	}

	for _, c := range cases {
		result := ShouldGoDiagonal(c.input.hp, c.input.tp)
		if result != c.expected {
			errStr := fmt.Sprintf("Fail. Expected: %v but got: %v\n", c.expected, result)
			errStr += fmt.Sprintf("HP: %v, TP: %v", c.input.hp, c.input.tp)
			t.Error(errStr)
		}
	}
}

func TestAreTheyTouching(t *testing.T) {
	cases := []struct {
		input struct {
			hp, tp *Pos
		}
		expected bool
	}{
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 0,
				y: 0,
			}, tp: &Pos{
				x: 0,
				y: 0,
			}},
			expected: true,
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 1,
				y: 0,
			}, tp: &Pos{
				x: 0,
				y: 0,
			}},
			expected: true,
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 1,
				y: 5,
			}, tp: &Pos{
				x: -1,
				y: 0,
			}},
			expected: false,
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 4,
				y: 4,
			}, tp: &Pos{
				x: 0,
				y: 0,
			}},
			expected: false,
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 1,
				y: 1,
			}, tp: &Pos{
				x: 0,
				y: 0,
			}},
			expected: true,
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 1,
				y: 2,
			}, tp: &Pos{
				x: 0,
				y: 0,
			}},
			expected: false,
		},
	}

	for _, c := range cases {
		result := AreTheyTouching(c.input.hp, c.input.tp)
		if result != c.expected {
			errStr := fmt.Sprintf("Fail. Expected: %v but got: %v\n", c.expected, result)
			errStr += fmt.Sprintf("HP: %v, TP: %v", c.input.hp, c.input.tp)
			t.Error(errStr)
		}
	}
}

func TestGetDirectionToMove(t *testing.T) {
	cases := []struct {
		input struct {
			hp, tp *Pos
		}
		expected []Dir
	}{
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 2,
				y: 2,
			}, tp: &Pos{
				x: 2,
				y: 0,
			}},
			expected: []Dir{
				{
					direction: "U",
				},
			},
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 2,
				y: 0,
			}, tp: &Pos{
				x: 2,
				y: 2,
			}},
			expected: []Dir{
				{
					direction: "D",
				},
			},
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 0,
				y: 2,
			}, tp: &Pos{
				x: 2,
				y: 2,
			}},
			expected: []Dir{
				{
					direction: "L",
				},
			},
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 2,
				y: 2,
			}, tp: &Pos{
				x: 0,
				y: 2,
			}},
			expected: []Dir{
				{
					direction: "R",
				},
			},
		},
		{
			input: struct{ hp, tp *Pos }{hp: &Pos{
				x: 2,
				y: 2,
			}, tp: &Pos{
				x: 0,
				y: 0,
			}},
			expected: []Dir{
				{
					direction: "R",
				},
				{
					direction: "U",
				},
			},
		},
	}

	for _, c := range cases {
		result := GetDirectionToMove(c.input.hp, c.input.tp)

		if len(result) != len(c.expected) {
			t.Error(fmt.Sprintf("Different direction length"))
		}

		if len(result) == 1 {
			if result[0].direction != c.expected[0].direction {
				errStr := fmt.Sprintf("Fail. Expected: %v but got: %v\n", c.expected, result)
				errStr += fmt.Sprintf("HP: %v, TP: %v\n", c.input.hp, c.input.tp)
				t.Error(errStr)
			}
		} else if len(result) == 2 {
			if result[0].direction != c.expected[0].direction && result[1].direction != c.expected[1].direction {
				errStr := fmt.Sprintf("Fail. Expected: %v but got: %v\n", c.expected, result)
				errStr += fmt.Sprintf("HP: %v, TP: %v\n", c.input.hp, c.input.tp)
				t.Error(errStr)
			}
		} else {
			t.Error("Too many directions")
		}
	}
}

func TestMoveTheRope_2Knots(t *testing.T) {
	cases := []struct {
		input    []string
		knot     int
		expected int
	}{
		{
			input: []string{
				"R 3",
			},
			knot:     2,
			expected: 3,
		},
		{
			input: []string{
				"R 12",
			},
			knot:     2,
			expected: 12,
		},
		{
			input: []string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
				"R 4",
				"D 1",
				"L 5",
				"R 2",
				"U 2",
			},
			knot:     2,
			expected: 14,
		},
		{
			input: []string{
				"U 6",
			},
			knot:     4,
			expected: 4,
		},
	}

	for _, c := range cases {
		var visitMap = map[string]int{
			"": 1,
		}
		var knots = GenerateKnots(c.knot)
		for _, v := range c.input {
			visitMap = MoveTheRope(v, visitMap, knots)
		}
		result := len(visitMap)
		if result != c.expected {
			errStr := fmt.Sprintf("Fail. Knot: %v. Expected: %v but got: %v\n", c.knot, c.expected, result)
			t.Error(errStr)
		}
	}
}
