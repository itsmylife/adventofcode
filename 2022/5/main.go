package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/itsmylife/adventofcode/2022/helper"
)

func main() {
	fileScanner, closeFile := helper.ReadFile("5")
	defer closeFile()

	row := 0
	cratesOnTop := ""
	initialCrates := [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if row < 8 {
			ReadInitialCrates(line, initialCrates)
		}

		if row == 9 {
			for _, v := range initialCrates {
				for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
					v[i], v[j] = v[j], v[i]
				}
			}
		}

		if row > 9 && line != "" {
			MoveCrates(line, initialCrates)
		}
		row++
	}

	cratesOnTop = ReadTopCrates(initialCrates)
	fmt.Println(cratesOnTop)
}

func ReadInitialCratesByLetter(line string, initialCrates [][]string) {
	line += " "
	ll := len(line)
	for i := 0; i < ll; i += 4 {
		sl := strings.Trim(line[i:i+4], " ")
		if sl != "" {
			initialCrates[i/4] = append(initialCrates[i/4], sl[1:2])
		}
	}
}

func ReadInitialCrates(line string, initialCrates [][]string) {
	line += " "
	ll := len(line)
	for i := 0; i < ll; i += 4 {
		sl := strings.Trim(line[i:i+4], " ")
		if sl != "" {
			initialCrates[i/4] = append(initialCrates[i/4], sl)
		}
	}
}

func MoveCrates(line string, crates [][]string) {
	words := strings.Split(line, " ")
	howMany, _ := strconv.ParseInt(words[1], 10, 32)
	from, _ := strconv.ParseInt(words[3], 10, 32)
	to, _ := strconv.ParseInt(words[5], 10, 32)

	hmi := int(howMany)
	fi := from - 1
	ti := to - 1

	cratesToMove := crates[fi][len(crates[fi])-hmi:]
	crates[fi] = crates[fi][0 : len(crates[fi])-hmi]
	for i := len(cratesToMove) - 1; i >= 0; i-- {
		crates[ti] = append(crates[ti], cratesToMove[i])
	}
}

func ReadTopCrates(crates [][]string) string {
	word := ""
	for _, v := range crates {
		word += v[len(v)-1][1:2]
	}
	return word
}
