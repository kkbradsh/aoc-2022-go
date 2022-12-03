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

var XYZ map[string]int = map[string]int{
	"X": LOSE,
	"Y": DRAW,
	"Z": WIN,
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
		myOutcome := XYZ[encodedPlays[1]]
		myPlay := ""
		for guess := range RUBRIC[theirPlay] {
			if RUBRIC[theirPlay][guess] == myOutcome {
				myPlay = guess
			}
		}
		score = score + SCORE[myPlay] + myOutcome
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
