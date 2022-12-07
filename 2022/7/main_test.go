package main

import (
	"fmt"
	"testing"
)

func TestDay7(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "1",
			expected: "1",
		},
	}

	for _, c := range cases {
		if c.input != c.expected {
			t.Error(fmt.Sprintf("Fail. Expected: %v but got: %v", c.expected, c.input))
		}
	}
}
