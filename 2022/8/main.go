package main

import (
	"fmt"
	"strconv"

	"github.com/itsmylife/adventofcode/2022/helper"
)

type Node struct {
	parent string
	pwd    string
	size   int64
}

func main() {
	fileScanner, closeFile := helper.ReadFile("8")
	defer closeFile()

	var treeMap []string

	li := 0
	var rows = make([]int64, 99)
	var cols = make([]int64, 99)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		rows, cols = CalculateHighest(line, li, rows, cols)
		treeMap = append(treeMap, line)
		li++
	}

	visibleTrees := FindVisibleTreeCount(treeMap)

	fmt.Println(fmt.Sprintf("Total visible tree count is : %d", visibleTrees))
}

func FindScenicScore() int {
	return 0
}

func CalculateHighest(line string, li int, rows, cols []int64) ([]int64, []int64) {
	var bigInRow int64
	for i, vv := range line {
		v := int64(vv) - 48
		if v > cols[i] {
			cols[i] = v
		}
		if v > bigInRow {
			bigInRow = v
		}
	}
	rows[li] = bigInRow

	return rows, cols
}

func FindVisibleTreeCount(input []string) int {

	// We assume cols and rows has same length
	ll := len(input[0])
	visible := ll*2 + (ll-2)*2

	for r := 1; r < ll-1; r++ {
		for c := 1; c < ll-1; c++ {
			t := ci(input[r][c])

			isVisible := true
			// Check in row

			for i := c - 1; i >= 0; i-- {
				rl := ci(input[r][i])
				if t <= rl {
					// it is invisible no need to check the rest
					isVisible = false
					break
				}
			}
			if isVisible {
				visible++
				continue
			}

			isVisible = true
			for i := c + 1; i < ll; i++ {
				rr := ci(input[r][i])
				if t <= rr {
					// it is invisible no need to check the rest
					isVisible = false
					break
				}
			}
			if isVisible {
				visible++
				continue
			}
			// Check in col

			isVisible = true
			for i := r - 1; i >= 0; i-- {
				ct := ci(input[i][c])
				if t <= ct {
					// it is invisible no need to check the rest
					isVisible = false
					break
				}
			}
			if isVisible {
				visible++
				continue
			}
			isVisible = true
			for i := r + 1; i < ll; i++ {
				cb := ci(input[i][c])
				if t <= cb {
					// it is invisible no need to check the rest
					isVisible = false
					break
				}
			}
			if isVisible {
				visible++
				continue
			}
		}
	}

	return visible
}

func ci(str uint8) int {
	ts, _ := strconv.ParseInt(string(str), 10, 64)
	return int(ts)
}
