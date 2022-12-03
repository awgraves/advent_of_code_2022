package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
Opponent
A = Rock
B = Paper
C = Scissors

Me
X = Rock
Y = Paper
Z = Scissors


Scoring
I. Shape selected
Rock = 1
Paper = 2
Scissors = 3

II. Outcome of round
Lost = 0
Tie = 3
Win = 6

Shape chosen (regardless if won) + outcome = score for round
*/

type scoreCounter struct {
	totalScore      int
	normalizeMap    map[string]string
	shapeScoreMap   map[string]int
	outcomeScoreMap map[string]int
	partTwoMoveMap  map[string]string
	winningMoveMap  map[string]string
}

func newScoreCounter() *scoreCounter {
	normalizeMap := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}
	shapeScoreMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	winningMoveMap := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}
	return &scoreCounter{
		normalizeMap:   normalizeMap,
		shapeScoreMap:  shapeScoreMap,
		winningMoveMap: winningMoveMap,
	}
}

func (sc *scoreCounter) scoreRoundPartOne(opp, me string) {
	me = sc.normalizeMap[me]
	sc.scoreRound(opp, me)
}

func (sc *scoreCounter) scoreRoundPartTwo(opp, outcome string) {
	myMove := sc.getMyMove(opp, outcome)
	sc.scoreRound(opp, myMove)
}

func (sc *scoreCounter) scoreRound(opp, me string) {
	sc.totalScore += sc.shapeScoreMap[me] + sc.getOutcomeScore(opp, me)
}

func (sc *scoreCounter) getMyMove(opp, outcome string) string {
	switch outcome {
	case "X": // lose
		var move string
		for l, w := range sc.winningMoveMap {
			if w == opp {
				move = l
				break
			}
		}
		return move
	case "Y": // tie
		return opp
	case "Z": // win
		return sc.winningMoveMap[opp]
	default:
		panic(outcome)
	}
}

func (sc *scoreCounter) getOutcomeScore(opp string, me string) int {
	if opp == me {
		return 3 // tie
	}
	winningMove := sc.winningMoveMap[opp]
	if me == winningMove {
		return 6
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	partOneCounter := newScoreCounter()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		partOneCounter.scoreRoundPartOne(moves[0], moves[1])
	}
	fmt.Printf("My final score (part 1): %d\n", partOneCounter.totalScore) // part 1 solution 8890

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	partTwoCounter := newScoreCounter()
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		partTwoCounter.scoreRoundPartTwo(moves[0], moves[1])
	}
	fmt.Printf("My final score (part 2): %d", partTwoCounter.totalScore) // part 2 solution 10238
}
