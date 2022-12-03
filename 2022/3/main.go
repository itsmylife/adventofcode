package main

import (
	"fmt"

	"github.com/itsmylife/adventofcode/2022/helper"
)

const (
	lowercaseOffset = 97 - 1
	uppercaseOffset = 65 - 27
)

func main() {
	fileScanner, closeFile := helper.ReadFile("3")
	defer closeFile()

	totalPoints := 0
	totalGroupPoints := 0
	var group []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		totalPoints += CalculatePriorityPoint(line)
		group = append(group, line)
		if len(group) == 3 {
			totalGroupPoints += CalculateGroupPriorityPoint(group)
			group = []string{}
		}
	}

	fmt.Println(totalPoints, totalGroupPoints)
}

func CalculateGroupPriorityPoint(group []string) int {
	for _, c1 := range group[0] {
		for _, c2 := range group[1] {
			if c1 == c2 {
				for _, c3 := range group[2] {
					if c2 == c3 {
						return getPoint(c1)
					}
				}
			}
		}
	}
	return 0
}

// CalculatePriorityPoint https://adventofcode.com/2022/day/3
func CalculatePriorityPoint(line string) int {
	ll := len(line)
	p1 := line[0 : ll/2]
	p2 := line[ll/2:]
	foundPoints := 0
	for _, c1 := range p1 {
		for _, c2 := range p2 {
			if c2 == c1 {
				return getPoint(c1)
			}
		}
	}
	return foundPoints
}

func getPoint(c rune) int {
	offset := lowercaseOffset
	if isUpperCase(c) {
		offset = uppercaseOffset
	}

	return int(c) - offset
}

func isUpperCase(c rune) bool {
	if int(c) < 97 {
		return true
	}
	return false
}
