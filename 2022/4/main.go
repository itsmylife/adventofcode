package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/itsmylife/adventofcode/2022/helper"
)

func main() {
	fileScanner, closeFile := helper.ReadFile("4")
	defer closeFile()

	fullyContainedPairsCount := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fullyContainedPairsCount += FindContainedRanges(line)
	}
	fmt.Println(fullyContainedPairsCount)
}

func FindContainedRanges(ranges string) int {
	elves := strings.Split(ranges, ",")
	e1 := strings.Split(elves[0], "-")
	e2 := strings.Split(elves[1], "-")

	e1s, _ := strconv.ParseInt(e1[0], 10, 32)
	e1e, _ := strconv.ParseInt(e1[1], 10, 32)
	e2s, _ := strconv.ParseInt(e2[0], 10, 32)
	e2e, _ := strconv.ParseInt(e2[1], 10, 32)

	if e1s <= e2s && e1e >= e2e {
		return 1
	} else if e1s >= e2s && e1e <= e2e {
		return 1
	}

	return 0
}
