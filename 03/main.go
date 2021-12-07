/*
	My initial hunch was to use os.Seek(?????) to loop through all
	first characters in every line and calculate the gamma and epsilon -- 
	just to avoid messing around with loops, but I couldn't figure out
	how to make that work so I ended up just using two for loops to compute the 
	gamma and epsilon and then convert those into int64s

	The error handling and file parsing is pretty tedious which is probably a 
	sign I should extract them into a util?

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	// initial gamma and epsilon are empty strings that we'll append to
	gamma, epsilon := "", ""

	s := bufio.NewScanner(file)
	var lines []string

	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}

	breadth, length := len(lines[0]), len(lines)

	// for every column in the diagnostic report, loop through every line
	// and count the number of zeroes. 
	for i := 0; i < breadth; i++ {
		zeroes := 0
		for j := 0; j < length; j++ {
			if lines[j][i] == '0' {
				zeroes++
			}
		}
		// if zeroes are greater than half the 'length' of the line, then we can conclude they
		// are the majority and append to gamma. Otherwise append to epsilon
		if zeroes > length/2 {
			gamma += "0"
			epsilon += "1"
		} else {
			epsilon += "0"
			gamma += "1"
		}
	}
	grate, _ := strconv.ParseInt(gamma, 2, 64)
	erate, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(grate * erate)
}
