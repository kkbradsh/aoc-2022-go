package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(fileName string) []string {
	// Load input file
	content, error := os.ReadFile(fileName)
	if error != nil {
		log.Fatal(error)
	}
	return strings.Split(string(content), "\r\n")
}

func ParseArrayGroupBySum(input []string) []int {
	output := []int{0}
	for _, line := range input {
		if string(line) == "" {
			output = append(output, 0)
		} else {
			value, _ := strconv.Atoi(line)
			output[len(output)-1] = output[len(output)-1] + value
		}
	}
	return output
}
