package main

import (
	"aoc/utils"
	"fmt"
	"sort"
)

func process(fileName string) int {
	input := utils.ReadFile(fileName)
	caloriesPerElf := utils.ParseArrayGroupBySum(input)
	// Return sum of 3 largest values
	sort.Ints(caloriesPerElf)
	sum := 0
	for i := 1; i <= 3; i++ {
		sum = sum + caloriesPerElf[len(caloriesPerElf)-i]
	}
	return sum
}

func main() {
	fmt.Println(process("final.txt"))
}
