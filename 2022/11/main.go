package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"

	"github.com/itsmylife/adventofcode/2022/helper"
)

type Monkey struct {
	items     []int
	throwMap  map[bool]int
	test      int
	operation *Operation
}

type Operation struct {
	input1 string
	op     string
	input2 string
}

func main() {
	fileScanner, closeFile := helper.ReadFile("11")
	defer closeFile()

	li := 1
	var monkeys []*Monkey

	for fileScanner.Scan() {
		line := fileScanner.Text()

		monkeys = ParseInput(line, li, monkeys)

		li++
	}

	biznis := RunTheSimulation(monkeys, 10000, false)
	sort.Ints(biznis)

	biz := biznis[7] * biznis[6]
	fmt.Println(fmt.Sprintf("Total Monkey Biz : %d", biz))
}

func RunTheSimulation(monkeys []*Monkey, rounds int, feelRelief bool) []int {
	round := 0
	mod := 1
	var biznis = make([]int, len(monkeys))

	for _, m := range monkeys {
		mod *= m.test
	}

	for round < rounds {
		for i, m := range monkeys {
			for _, it := range m.items {

				var i1, i2, l int
				if m.operation.input1 == "old" {
					i1 = it
				}
				if m.operation.input2 == "old" {
					i2 = it
				} else {
					i2 = helper.ConvertInt(m.operation.input2)
				}

				switch m.operation.op {
				case "+":
					l = i1 + i2
				case "*":
					l = i1 * i2
				default:
					log.Fatalln(fmt.Sprintf("Operation [%s] does not implemented", m.operation.op))
				}

				if feelRelief {
					l = int(math.Floor(float64(l) / 3))
				} else {
					l %= mod
				}

				nextMonkeyIndex := -1
				if l%m.test == 0 {
					nextMonkeyIndex = m.throwMap[true]
				} else {
					nextMonkeyIndex = m.throwMap[false]
				}

				monkeys[nextMonkeyIndex].items = append(monkeys[nextMonkeyIndex].items, l)
			}
			biznis[i] += len(m.items)
			m.items = []int{}
		}
		round++
	}

	return biznis
}

func ParseInput(line string, li int, monkeys []*Monkey) []*Monkey {
	lineNumber := math.Mod(float64(li), 7)
	monkeyIndex := int(math.Floor(float64(li / 7)))

	if len(monkeys) <= monkeyIndex {
		monkeys = append(monkeys, &Monkey{
			items: []int{},
			throwMap: map[bool]int{
				true:  -1,
				false: -1,
			},
			test: 0,
			operation: &Operation{
				input1: "",
				op:     "",
				input2: "",
			},
		})
	}

	switch lineNumber {
	case 2:
		parts := strings.Split(line, ": ")
		nums := strings.Split(parts[1], ", ")
		for _, v := range nums {
			monkeys[monkeyIndex].items = append(monkeys[monkeyIndex].items, helper.ConvertInt(v))
		}

	case 3:
		parts := strings.Split(line, " = ")
		parts = strings.Split(parts[1], " ")
		monkeys[monkeyIndex].operation = &Operation{
			input1: parts[0],
			op:     parts[1],
			input2: parts[2],
		}
	case 4:
		parts := strings.Split(line, " ")
		monkeys[monkeyIndex].test = helper.ConvertInt(parts[len(parts)-1])
	case 5:
		parts := strings.Split(line, " ")
		monkeys[monkeyIndex].throwMap = map[bool]int{
			true: helper.ConvertInt(parts[len(parts)-1]),
		}
	case 6:
		parts := strings.Split(line, " ")
		monkeys[monkeyIndex].throwMap[false] = helper.ConvertInt(parts[len(parts)-1])
	}

	return monkeys
}
