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
	overlappingPairsCount := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fullyContainedPairsCount += FindContainedRanges(line)
		overlappingPairsCount += FindOverlappingRanges(line)
	}
	fmt.Println(fullyContainedPairsCount, overlappingPairsCount)
}

func FindContainedRanges(ranges string) int {
	e1s, e1e, e2s, e2e := splitToElves(ranges)

	if e1s <= e2s && e1e >= e2e {
		return 1
	} else if e1s >= e2s && e1e <= e2e {
		return 1
	}

	return 0
}

func FindOverlappingRanges(ranges string) int {
	e1s, e1e, e2s, e2e := splitToElves(ranges)

	if e1s < e2s && e1e < e2s {
		return 0
	} else if e1s > e2s && e2e < e1s {
		return 0
	}

	return 1
}

func splitToElves(ranges string) (e1s, e1e, e2s, e2e int64) {
	elves := strings.Split(ranges, ",")
	e1 := strings.Split(elves[0], "-")
	e2 := strings.Split(elves[1], "-")

	e1s, _ = strconv.ParseInt(e1[0], 10, 32)
	e1e, _ = strconv.ParseInt(e1[1], 10, 32)
	e2s, _ = strconv.ParseInt(e2[0], 10, 32)
	e2e, _ = strconv.ParseInt(e2[1], 10, 32)
	// return e1s, e1e, e2s, e2e
	return
}
