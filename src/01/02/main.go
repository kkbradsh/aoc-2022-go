package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

	// Find 3 elves with most calories
	max := []int{}
	for index, value := range elves {
		if index < 3 {
			max = append(max, value)
		} else {
			// Sort to locate the smallest number in max
			sort.Ints(max)
			// Replace if new value is greater than smallest value in max
			if value > max[0] {
				max[0] = value
			}
		}
	}

	// Return sum of top 3
	sum := 0
	for _, value := range max {
		sum = sum + value
	}
	return sum
}

func main() {
	fmt.Println(process("final.txt"))
}
