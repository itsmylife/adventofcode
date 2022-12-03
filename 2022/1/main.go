package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/itsmylife/adventofcode/2022/helper"
)

// https://adventofcode.com/2022/day/1
func main() {
	fileScanner, closeFile := helper.ReadFile("1")
	defer closeFile()

	elfIndex := 0
	maxCalorie := 0
	maxCalorieIndex := 0
	currentCalorie := 0
	var calorieList []int

	for fileScanner.Scan() {
		txt := fileScanner.Text()

		if txt == "" {
			currentCalorie = 0
			elfIndex += 1
		} else {
			cal, err := strconv.Atoi(txt)
			if err != nil {
				log.Fatalln(err)
			}

			currentCalorie += cal

			if currentCalorie > maxCalorie {
				maxCalorie = currentCalorie
				maxCalorieIndex = elfIndex
			}

			if len(calorieList) < elfIndex+1 {
				calorieList = append(calorieList, currentCalorie)
			} else {
				calorieList[elfIndex] = currentCalorie
			}
		}
	}

	totalOfTopThree := 0
	sort.Sort(sort.Reverse(sort.IntSlice(calorieList)))
	for _, v := range calorieList[:3] {
		totalOfTopThree += v
	}

	fmt.Println(fmt.Sprintf("Max Calorie: %d --- Index: %d", maxCalorie, maxCalorieIndex))
	fmt.Println(fmt.Sprintf("Total Of Top Three: %d", totalOfTopThree))
}
