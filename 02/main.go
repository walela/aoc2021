package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Pretty similar to day 1 and pretty straightforward again.
// I decided to do the computation in one pass this time which
// made things a little neater. I'm not doing much error handling
// though, which is something I might decide to do in future
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
		delta, _ := strconv.Atoi(command[1])
		switch command[0] {
		case "forward":
			alignment += delta
		case "up":
			depth -= delta
		case "down":
			depth += delta
		default:
		}
	}
	fmt.Println(depth * alignment)
}
