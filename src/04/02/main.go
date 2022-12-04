package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func process(fileName string) int {
	input := utils.ReadFile(fileName)
	// Process input
	count := 0
	for i := 0; i < len(input); i++ {
		// Get sections
		sections := strings.Split(input[i], ",")
		section1 := strings.Split(sections[0], "-")
		section2 := strings.Split(sections[1], "-")
		a, _ := strconv.Atoi(section1[0])
		b, _ := strconv.Atoi(section1[1])
		c, _ := strconv.Atoi(section2[0])
		d, _ := strconv.Atoi(section2[1])
		// Contained?
		if (a <= c && b >= c) ||
			(a <= d && b >= d) ||
			(c <= a && d >= a) ||
			(c <= b && d >= b) {
			count = count + 1
		}
	}
	return count
}

func main() {
	fmt.Println(process("final.txt"))
}
