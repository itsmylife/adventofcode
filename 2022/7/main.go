package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/itsmylife/adventofcode/2022/helper"
)

type Node struct {
	parent string
	pwd    string
	size   int64
}

func main() {
	fileScanner, closeFile := helper.ReadFile("7")
	defer closeFile()

	var dirMap = make(map[string]Node)
	pwd := ""
	lineNum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		lineNum++
		fmt.Println(fmt.Sprintf("Line read: %d", lineNum))

		parts := strings.Split(line, " ")

		if parts[0] == "$" {
			// Out of the directory
			if parts[1] == "cd" && parts[2] == ".." {
				fmt.Println("go one level up")
				if currentNode, ok := dirMap[pwd]; ok {
					pwd = currentNode.parent
				} else {
					log.Fatalln("No such directory!!!!!!!!")
				}
			} else
			// Into the directory
			if parts[1] == "cd" && parts[2] != ".." {
				newPwd := ""
				if parts[2] == "/" {
					newPwd = parts[2]
				} else if pwd == "/" {
					newPwd = pwd + parts[2]
				} else {
					newPwd = pwd + "/" + parts[2]
				}
				if _, ok := dirMap[newPwd]; !ok {
					dirMap[newPwd] = Node{
						parent: pwd,
						pwd:    newPwd,
						size:   0,
					}
				}
				pwd = newPwd
			}
		} else {
			if parts[0] != "dir" {
				fs, _ := strconv.ParseInt(parts[0], 10, 64)
				UpdateSize(fs, pwd, dirMap)
			}
		}
	}

	var total int64 = 0
	for k, v := range dirMap {
		if v.size < 100000 {
			fmt.Println(k)
			total += v.size
		}
	}

	fmt.Println(fmt.Sprintf("The sum is: %d", total))
}

func UpdateSize(fs int64, pwd string, dirMap map[string]Node) {
	if node, ok := dirMap[pwd]; ok {
		node.size += fs
		dirMap[pwd] = node
		if node.parent != "" {
			UpdateSize(fs, node.parent, dirMap)
		}
	}
}
