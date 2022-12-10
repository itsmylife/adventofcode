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

type CRT struct {
	spritePos int
	cycle     int
	rowIndex  int
	rows      []string
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

	crt := &CRT{
		spritePos: 1,
		cycle:     0,
		rowIndex:  0,
		rows:      []string{"", "", "", "", "", ""},
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		op, val := Read(line)

		execution.operation = Operation{
			name: op,
			val:  val,
		}
		Execute(execution)
		DisplayOnCrt(crt, Operation{
			name: op,
			val:  val,
		})
	}

	fmt.Println(fmt.Sprintf("Total strebgth : %d", execution.strength))
	for _, v := range crt.rows {
		fmt.Println(v)
	}
}

func DisplayOnCrt(crt *CRT, op Operation) {
	if op.name == "noop" {
		DrawOnCrt(crt)
		crt.cycle++
	} else if op.name == "addx" {
		DrawOnCrt(crt)
		crt.cycle++

		DrawOnCrt(crt)
		crt.cycle++
		crt.spritePos += op.val
	}
}

func DrawOnCrt(crt *CRT) {
	if crt.cycle == 40 {
		crt.cycle = 0
		crt.rowIndex++
	}
	if crt.rowIndex == 6 {
		return
	}
	if crt.cycle >= crt.spritePos-1 && crt.cycle <= crt.spritePos+1 {
		// put a #
		crt.rows[crt.rowIndex] += "#"
	} else {
		// put a .
		crt.rows[crt.rowIndex] += "."
	}
}

// Part 1 -------------------------------------------

func Execute(ex *Execution) {
	if ex.operation.name == "noop" {
		ex.cycle++
		CheckPoint(ex)
	} else if ex.operation.name == "addx" {
		ex.cycle++
		CheckPoint(ex)

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
