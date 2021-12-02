package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	depth := 0
	alignment := 0
	s := bufio.NewScanner(file)

	for s.Scan() {
		command := strings.Split(s.Text(), " ")
		switch command[0] {
		case "forward":
			delta, _ := strconv.Atoi(command[1])
			alignment += delta
		case "up":
			delta, _ := strconv.Atoi(command[1])
			depth -= delta
		case "down":
			delta, _ := strconv.Atoi(command[1])
			depth += delta
		default:
		}
	}
	fmt.Println(depth * alignment)
}
