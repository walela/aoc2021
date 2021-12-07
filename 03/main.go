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

	// gamma, epsilon := 0, 0

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
		fmt.Println(zeroes)
	}
}
