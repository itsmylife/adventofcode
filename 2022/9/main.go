package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/itsmylife/adventofcode/2022/helper"
)

type Pos struct {
	x int
	y int
}

type Dir struct {
	direction string
}

func main() {
	fileScanner, closeFile := helper.ReadFile("9")
	defer closeFile()

	var visitMap = map[string]int{
		"": 1,
	}

	li := 0
	headPos := &Pos{
		x: 0,
		y: 0,
	}
	tailPos := &Pos{
		x: 0,
		y: 0,
	}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		visitMap = MoveOnTheRope(line, visitMap, headPos, tailPos, li)
		li++
	}

	visits := len(visitMap)

	fmt.Println(fmt.Sprintf("Total visits : %d", visits))
}

func MoveOnTheRope(line string, vm map[string]int, hp, tp *Pos, li int) map[string]int {
	ps := strings.Split(line, " ")
	c := ci(ps[1])
	for i := 0; i < c; i++ {
		Move(Dir{
			direction: ps[0],
		}, hp)
		pos := MoveTail(hp, tp)
		if _, ok := vm[pos]; !ok {
			vm[pos] = 1
		}
	}

	return vm
}

func Move(dir Dir, hp *Pos) {
	switch dir.direction {
	case "U":
		hp.y += 1
	case "D":
		hp.y -= 1
	case "R":
		hp.x += 1
	case "L":
		hp.x -= 1
	}
}

func MoveTail(hp *Pos, tp *Pos) string {
	if AreTheyTouching(hp, tp) {
		return ""
	}

	dirs := GetDirectionToMove(hp, tp)
	for _, v := range dirs {
		Move(v, tp)
	}

	return fmt.Sprintf("%d,%d", tp.x, tp.y)
}

func GetDirectionToMove(hp *Pos, tp *Pos) []Dir {
	var dirs []Dir
	if ShouldGoDiagonal(hp, tp) {
		if tp.x > hp.x {
			// go left
			dirs = append(dirs, Dir{
				direction: "L",
			})
		}
		if tp.x < hp.x {
			// go right
			dirs = append(dirs, Dir{
				direction: "R",
			})
		}
		if tp.y > hp.y {
			// go down
			dirs = append(dirs, Dir{
				direction: "D",
			})
		}
		if tp.y < hp.y {
			// go up
			dirs = append(dirs, Dir{
				direction: "U",
			})
		}
	} else {
		if tp.x-hp.x >= 2 {
			// go left
			dirs = append(dirs, Dir{
				direction: "L",
			})
		} else if tp.x-hp.x <= -2 {
			// go right
			dirs = append(dirs, Dir{
				direction: "R",
			})
		} else if tp.y-hp.y >= 2 {
			// go down
			dirs = append(dirs, Dir{
				direction: "D",
			})
		} else if tp.y-hp.y <= -2 {
			// go up
			dirs = append(dirs, Dir{
				direction: "U",
			})
		}
	}

	return dirs
}

func ShouldGoDiagonal(hp *Pos, tp *Pos) bool {
	if math.Abs(float64(hp.x-tp.x)) < 2 && math.Abs(float64(hp.y-tp.y)) < 2 {
		return false
	}
	return true
}

func AreTheyTouching(hp *Pos, tp *Pos) bool {
	if hp.x == tp.x && hp.y == tp.y {
		return true
	}

	if math.Abs(float64(hp.x-tp.x)) >= 2 || math.Abs(float64(hp.y-tp.y)) >= 2 {
		return false
	}

	return true
}

func ci(str string) int {
	ts, _ := strconv.ParseInt(str, 10, 64)
	return int(ts)
}
