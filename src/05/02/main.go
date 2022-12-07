package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func processStacks(stacks [][]byte, instructions [][]string) [][]byte {
	// Process instructions
	for _, element := range instructions {
		move, _ := strconv.Atoi(element[0])
		fromStack, _ := strconv.Atoi(element[1])
		toStack, _ := strconv.Atoi(element[2])
		fromIndex := fromStack - 1
		toIndex := toStack - 1
		newStack := make([]byte, len(stacks[toIndex])+move)
		for i := 0; i < move; i++ {
			newStack[i] = stacks[fromIndex][0]
			stacks[fromIndex] = stacks[fromIndex][1:]
		}
		for index, element := range stacks[toIndex] {
			newStack[index+move] = element
		}
		stacks[toIndex] = newStack
	}
	return stacks
}

func parseInstruction(instructions [][]string, line string) [][]string {
	line = strings.Replace(line, "move ", "", 1)
	line = strings.Replace(line, "from ", "", 1)
	line = strings.Replace(line, "to ", "", 1)
	values := strings.Split(line, " ")
	return append(instructions, values)
}

func parseStack(stacks [][]byte, line string) [][]byte {
	SECTION_LENGTH := 4
	// first time through
	if len(stacks) == 0 {
		for i := 0; i < len(line); i = i + SECTION_LENGTH {
			stacks = append(stacks, []byte{})
		}
	}
	// each time through
	for i := 0; i < len(line); i = i + SECTION_LENGTH {
		crate := line[i+1]
		if crate >= 'A' && crate <= 'Z' {
			stacks[i/SECTION_LENGTH] = append(stacks[i/SECTION_LENGTH], crate)
		}
	}
	return stacks
}

type parseInputOutput struct {
	stacks       [][]byte
	instructions [][]string
}

func parseInput(input []string) parseInputOutput {
	stacks := [][]byte{}
	instructions := [][]string{}
	foundSeparator := false
	for _, line := range input {
		if line == "" {
			foundSeparator = true
		} else if !foundSeparator {
			stacks = parseStack(stacks, line)
		} else {
			instructions = parseInstruction(instructions, line)
		}
	}
	return parseInputOutput{stacks, instructions}
}

func process(fileName string) string {
	input := utils.ReadFile(fileName)
	parsed_input := parseInput(input)
	stacks := parsed_input.stacks
	instructions := parsed_input.instructions
	finalStacks := processStacks(stacks, instructions)
	// return top of the stacks
	top := ""
	for _, element := range finalStacks {
		top = top + string(element[0])
	}
	return top
}

func main() {
	fmt.Println(process("final.txt"))
}
