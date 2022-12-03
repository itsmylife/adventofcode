package helper

import (
	"bufio"
	"log"
	"os"
	"runtime"
	"strings"
)

func ReadFile(index string) (*bufio.Scanner, func() error) {
	mydir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	readFile, err := os.Open(mydir + "/2022/inputs/input-" + index + ".txt")
	if err != nil {
		log.Fatalln(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile.Close
}

func File() (fn string) {
	_, fn, _, _ = runtime.Caller(0)
	return
}

func Folder() string {
	path := File()
	str := strings.Split(path, "main.go")
	return str[0]
}

// IndexOf => Instead of this bytes.IndexByte() could be used
func IndexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
