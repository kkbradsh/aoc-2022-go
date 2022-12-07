package main

import (
	"aoc/utils"
	"fmt"
)

func findMarker(line string) int {
	MARKER_LENGTH := 4
	for i := MARKER_LENGTH; i < len(line); i++ {
		marker := []byte(line[i-MARKER_LENGTH : i])
		occured := map[byte]bool{}
		foundDup := false
		for _, element := range marker {
			if !occured[element] {
				occured[element] = true
			} else {
				foundDup = true
			}
		}
		if !foundDup {
			return i
		}
	}
	return 0
}

func process(fileName string) int {
	input := utils.ReadFile(fileName)
	return findMarker(input[0])
}

func main() {
	fmt.Println(process("final.txt"))
}
