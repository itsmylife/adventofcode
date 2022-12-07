package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/itsmylife/adventofcode/2022/helper"
)

type Node struct {
	parent string
	pwd    string
	size   int64
}

const totalSpace int64 = 70000000
const minRequiredSpace int64 = 30000000

func main() {
	fileScanner, closeFile := helper.ReadFile("7")
	defer closeFile()

	pwd := ""
	lineNum := 0
	var dirMap = make(map[string]Node)

	var sizeToDelete int64 = math.MaxInt64

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

	var minSpaceToDelete int64 = 0
	if rootNode, ok := dirMap["/"]; ok {
		fmt.Println(fmt.Sprintf("free space left: %d", totalSpace-rootNode.size))
		freeSpace := totalSpace - rootNode.size
		minSpaceToDelete = minRequiredSpace - freeSpace
	}

	var total int64 = 0
	for k, v := range dirMap {
		if v.size < 100000 {
			fmt.Println(k)
			total += v.size
		}

		if v.size > minSpaceToDelete && v.size < sizeToDelete {
			sizeToDelete = v.size
		}
	}

	fmt.Println(fmt.Sprintf("The sum is: %d", total))
	fmt.Println(fmt.Sprintf("Minimum to delete: %d", sizeToDelete))
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
