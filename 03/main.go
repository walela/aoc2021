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

	s := bufio.NewScanner(file)
	var lines []string

	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}
	fmt.Println(lines)
}
