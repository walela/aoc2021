package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// This is my very first foray into Go after a long period so I'm fairly confident
// a lot of this is a hacky mess.

// Either way, I decided to read the contents of the input file into an array and convert them
// into ints, and do a for loop comparing the current value with the previous value. Not sure if
// that's the fastest way as well!

// I'll probably revise this over time but pretty happy I got the solution at all...onwards :)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	var depths []int

	// loop through the file and append depth measurements to `depths` array
	for s.Scan() {
		depth, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		depths = append(depths, depth)
	}

	// loop through array of depths and increase count if there is an increase in depth between
	// two consecutive measurements
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}
	fmt.Println(count)
}
