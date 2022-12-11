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

	biznis := RunTheSimulation(monkeys)
	sort.Ints(biznis)

	biz := biznis[7] * biznis[6]
	fmt.Println(fmt.Sprintf("Total Monkey Biz : %d", biz))
}

func RunTheSimulation(monkeys []*Monkey) []int {
	round := 1
	var biznis = make([]int, len(monkeys))
	for round < 21 {
		for i, m := range monkeys {
			for _, it := range m.items {
				worryLevel := CalculateTheWorryLevel(m.operation, it)
				nextMonkeyIndex := DetermineNextMonkey(m.throwMap, m.test, worryLevel)
				if nextMonkeyIndex == i {
					fmt.Println("#######################################")
					fmt.Println("!!!!!Monkey is throwing it to itself!!!!!!")
					fmt.Println("#######################################")
				}
				monkeys[nextMonkeyIndex].items = append(monkeys[nextMonkeyIndex].items, worryLevel)
			}
			biznis[i] += len(m.items)
			m.items = []int{}
		}
		fmt.Println(fmt.Sprintf("Round %d completed.", round))
		for k := 0; k < len(monkeys); k++ {
			fmt.Println(fmt.Sprintf("Monkey %d: %v", k, monkeys[k].items))
		}
		fmt.Println("----------------------------------------")
		round++
	}

	return biznis
}

func CalculateTheWorryLevel(operation *Operation, it int) int {
	var i1, i2, l int
	if operation.input1 == "old" {
		i1 = it
	}
	if operation.input2 == "old" {
		i2 = it
	} else {
		i2 = helper.ConvertInt(operation.input2)
	}
	switch operation.op {
	case "+":
		l = i1 + i2
	case "*":
		l = i1 * i2
	default:
		log.Fatalln(fmt.Sprintf("Operation [%s] does not implemented", operation.op))
	}
	return int(math.Floor(float64(l) / 3))
}

func DetermineNextMonkey(throwMap map[bool]int, test int, level int) int {
	isDivisible := math.Mod(float64(level), float64(test))
	if isDivisible == 0 {
		return throwMap[true]
	}
	return throwMap[false]
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
	// case 1:
	// 	parts := strings.Split(line, " ")
	// 	if _, ok := monkeys[parts[1]]; !ok {
	// 		monkeys[parts[1]] = &Monkey{}
	// 	}
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
