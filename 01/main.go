package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	var depths []int
	for s.Scan() {
		depth, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		depths = append(depths, depth)
	}
	count := 0
	for index, element := range depths {
		if index > 0 {
			if element > depths[index-1] {
				fmt.Println(element, "is greater than", depths[index-1])
				count++
			}
		}

	}
	fmt.Println(count)
}
