package main

import (
	"fmt"
	"strings"

	"github.com/itsmylife/adventofcode/2022/helper"
)

type Operation struct {
	name string
	val  int
}

type Execution struct {
	registerX       int
	operation       Operation
	cycle           int
	multiplierIndex int
	strength        int
	strengths       []int
	line            string
}

var Multipliers = []int{20, 60, 100, 140, 180, 220}

func main() {
	fileScanner, closeFile := helper.ReadFile("10")
	defer closeFile()

	execution := &Execution{
		registerX: 1,
		operation: Operation{
			name: "",
			val:  0,
		},
		cycle:           0,
		multiplierIndex: 0,
		strength:        0,
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		op, val := Read(line)
		execution.operation = Operation{
			name: op,
			val:  val,
		}
		Execute(execution)
		fmt.Println(fmt.Sprintf("Cycle %v executed!", execution.cycle))
	}

	fmt.Println(fmt.Sprintf("Total strebgth : %d", execution.strength))
}

func Execute(ex *Execution) {
	if ex.operation.name == "noop" {
		ex.cycle++
		CheckPoint(ex)
	} else if ex.operation.name == "addx" {
		ex.cycle++
		CheckPoint(ex) // middle check

		ex.cycle++
		CheckPoint(ex)
		ex.registerX += ex.operation.val
	}
}

func CheckPoint(ex *Execution) {
	if ex.multiplierIndex < len(Multipliers) && ex.cycle == Multipliers[ex.multiplierIndex] {
		fmt.Println(fmt.Sprintf("Cycle: %d -- RegisterX: %d", ex.cycle, ex.registerX))
		fmt.Println(fmt.Sprintf("Multiplier Index: %d", ex.multiplierIndex))
		ex.strength += Multipliers[ex.multiplierIndex] * ex.registerX
		ex.strengths = append(ex.strengths, Multipliers[ex.multiplierIndex]*ex.registerX)
		ex.multiplierIndex++
	}
}

func Read(line string) (string, int) {
	parts := strings.Split(line, " ")

	if len(parts) == 1 {
		// noop operation
		return parts[0], 0
	}

	val := helper.ConvertInt(parts[1])
	return parts[0], val
}
