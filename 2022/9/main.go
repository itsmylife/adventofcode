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

	knot2 := GenerateKnots(2)
	knot10 := GenerateKnots(10)
	var visitMap2Knots = map[string]int{
		"": 1,
	}
	var visitMap10Knots = map[string]int{
		"": 1,
	}

	li := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		visitMap2Knots = MoveTheRope(line, visitMap2Knots, knot2)
		visitMap10Knots = MoveTheRope(line, visitMap10Knots, knot10)
		li++
	}

	shortKnotTail := len(visitMap2Knots)
	longKnotTail := len(visitMap10Knots)

	fmt.Println(fmt.Sprintf("Total shortKnotTail : %d", shortKnotTail))
	fmt.Println(fmt.Sprintf("Total longKnotTail : %d", longKnotTail))
}

func GenerateKnots(count int) []*Pos {
	var knots []*Pos
	for i := 0; i < count; i++ {
		knots = append(knots, &Pos{
			x: 0,
			y: 0,
		})
	}
	return knots
}

func MoveTheRope(line string, vm map[string]int, kn []*Pos) map[string]int {
	ps := strings.Split(line, " ")
	direction := ps[0]
	count := ci(ps[1])
	kl := len(kn)
	for i := 0; i < count; i++ {
		for j := 1; j < kl; j++ {
			Move(Dir{
				direction: direction,
			}, kn[j-1])

			for k := j; k < kl; k++ {
				pos := MoveTail(kn[k-1], kn[k])
				if k == kl-1 {
					if _, ok := vm[pos]; !ok {
						vm[pos] = 1
					}
				}
			}
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
