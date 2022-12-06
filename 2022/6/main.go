package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	mydir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	readFile, err := os.Open(mydir + "/2022/inputs/input-6.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer readFile.Close()

	const maxSz = 1
	b := make([]byte, maxSz)

	token := ""
	ri := 0

	for {
		readTotal, err := readFile.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		char := string(b[:readTotal])
		ri++

		u, m := FindMarker(char, token, 14)
		token = m

		if u {
			break
		}
	}

	fmt.Println(fmt.Sprintf("Last index read: %d -- and the marker is %s", ri, token))
}

func FindMarker(c string, token string, charLength int) (unique bool, marker string) {

	if len(token) == 0 {
		return false, c
	}

	for i, v := range token {
		if string(v) == c {
			token += c
			return false, token[i+1:]
		}
	}

	token += c

	if len(token) == charLength {
		return true, token
	}

	return false, token
}
