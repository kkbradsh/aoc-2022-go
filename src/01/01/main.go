package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func process(fileName string) int {
	// Load input file
	content, error := os.ReadFile(fileName)
	if error != nil {
		log.Fatal(error)
	}
	input := strings.Split(string(content), "\r\n")

	// Walk input and sum item calories by elves
	elves := []int{0}
	for _, line := range input {
		if string(line) == "" {
			elves = append(elves, 0)
		} else {
			value, _ := strconv.Atoi(line)
			elves[len(elves)-1] = elves[len(elves)-1] + value
		}
	}

	// find max calories
	max := 0
	for _, value := range elves {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	fmt.Println(process("final.txt"))
}
