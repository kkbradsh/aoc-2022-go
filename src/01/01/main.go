package main

import (
	"fmt"
	"aoc/utils"
)

func max(arr [] int) int {
	max := 0
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func process(fileName string) int {
	input := utils.ReadFile(fileName)
	caloriesPerElf := utils.ParseArrayGroupBySum(input)
	return max(caloriesPerElf)
}

func main() {
	fmt.Println(process("final.txt"))
}
