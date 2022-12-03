package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func sumPriorities(input []string) int {
	priorities := []string{}
	for i := 0; i < len(input); i = i + 3 {
		// Get elves
		compartment1 := input[i+0]
		compartment2 := input[i+1]
		compartment3 := input[i+2]
		// Find priorities
		common12 := ""
		for j := 0; j < len(compartment1); j++ {
			char := compartment1[j : j+1]
			if strings.Contains(compartment2, char) &&
				!strings.Contains(common12, char) {
				common12 = common12 + char
			}
		}
		common := ""
		for j := 0; j < len(common12); j++ {
			char := common12[j : j+1]
			if strings.Contains(compartment3, char) && !strings.Contains(common, char) {
				common = common + char
			}
		}
		priorities = append(priorities, common)
	}
	// Sum priorities
	// ascii value of 'A' is 65, priorities of A-Z is 27 - 52
	// ascii value of 'a' is 97, a-z is 1 - 26
	sum := 0
	for _, item := range priorities {
		ascii := []byte(item)
		if item == strings.ToLower(item) {
			sum = sum + int(ascii[0]) - 97 + 1
		} else {
			sum = sum + int(ascii[0]) - 65 + 27
		}
	}
	return sum
}

func process(fileName string) int {
	input := utils.ReadFile(fileName)
	return sumPriorities(input)
}

func main() {
	fmt.Println(process("final.txt"))
}
