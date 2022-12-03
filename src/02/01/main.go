package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

const ROCK string = "ROCK"
const PAPER string = "PAPER"
const SCISSORS string = "SCISSORS"

const LOSE int = 0
const DRAW int = 3
const WIN int = 6

var ABC map[string]string = map[string]string{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var XYZ map[string]string = map[string]string{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

var SCORE map[string]int = map[string]int{
	ROCK:     1,
	PAPER:    2,
	SCISSORS: 3,
}

var RUBRIC map[string]map[string]int = map[string]map[string]int{
	ROCK: {
		ROCK:     DRAW,
		PAPER:    WIN,
		SCISSORS: LOSE,
	},
	PAPER: {
		ROCK:     LOSE,
		PAPER:    DRAW,
		SCISSORS: WIN,
	},
	SCISSORS: {
		ROCK:     WIN,
		PAPER:    LOSE,
		SCISSORS: DRAW,
	},
}

func playGame(input []string) int {
	score := 0
	for i := 0; i < len(input); i++ {
		encodedPlays := strings.Split(input[i], " ")
		theirPlay := ABC[encodedPlays[0]]
		myPlay := XYZ[encodedPlays[1]]
		score = score + SCORE[myPlay] + RUBRIC[theirPlay][myPlay]
	}
	return score
}

func process(fileName string) int {
	input := utils.ReadFile(fileName)
	return playGame(input)
}

func main() {
	fmt.Println(process("final.txt"))
}
