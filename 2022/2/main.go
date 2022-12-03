package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/itsmylife/adventofcode/2022/helper"
)

// 0 A X Rock		Tas
// 1 B Y Paper		Kagit
// 2 C Z Scissors	Makas

var abc = []string{"A", "B", "C"}
var xyz = []string{"X", "Y", "Z"}

func main() {
	fileScanner, closeFile := helper.ReadFile("2")
	defer closeFile()
	totalPointsPart1 := 0
	totalPointsPart2 := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		cols := strings.Split(line, " ")
		index1 := helper.IndexOf(cols[0], abc)
		index2 := helper.IndexOf(cols[1], xyz)

		totalPointsPart1 += CalculateRockPaperScissors(index1, index2)
		totalPointsPart2 += CalculateWhatsToPlay(cols)
	}

	fmt.Println(fmt.Sprintf("Total Part1: %d", totalPointsPart1))
	fmt.Println(fmt.Sprintf("Total Part2: %d", totalPointsPart2))
}

// CalculateRockPaperScissors https://adventofcode.com/2022/day/2#part1
func CalculateRockPaperScissors(index1, index2 int) int {
	point := 0
	point += index2 + 1

	if index2 == index1 {
		// fmt.Println("Draw +3")
		point += 3
	} else if mod3(index1+1) == index2 {
		// fmt.Println("Won +6")
		point += 6
	} else if mod3(index1+2) == index2 {
		// fmt.Println("Lost 0")
	}

	return point
}

// CalculateWhatsToPlay https://adventofcode.com/2022/day/2#part2
func CalculateWhatsToPlay(cols []string) int {
	resultMap := map[string]int{
		"Y": 0,
		"Z": 1,
		"X": 2,
	}
	index1 := helper.IndexOf(cols[0], abc)
	index2 := mod3(index1 + resultMap[cols[1]])
	return CalculateRockPaperScissors(index1, index2)
}

func mod3(num int) int {
	return int(math.Mod(float64(num), 3))
}
