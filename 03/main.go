package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	// TODO - research on different ways to hanlde errors (panic, exit codes etc)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	gamma, epsilon := "", ""

	s := bufio.NewScanner(file)
	var lines []string

	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}

	// i and j represent the breadth and length of the lines
	// matrix
	breadth, length := len(lines[0]), len(lines)
	for i := 0; i < breadth; i++ {
		zeroes := 0
		for j := 0; j < length; j++ {
			if lines[j][i] == '0' {
				zeroes++
			}
		}
		// if zeroes are greater than half, append to gamma
		if zeroes > length/2 {
			gamma += "0"
			epsilon += "1"
		} else {
			epsilon += "0"
			gamma += "1"
		}
	}
	fmt.Println(gamma, epsilon)
}
