package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

/*
[H]                 [Z]         [J]
[L]     [W] [B]     [G]         [R]
[R]     [G] [S]     [J] [H]     [Q]
[F]     [N] [T] [J] [P] [R]     [F]
[B]     [C] [M] [R] [Q] [F] [G] [P]
[C] [D] [F] [D] [D] [D] [T] [M] [G]
[J] [C] [J] [J] [C] [L] [Z] [V] [B]
[M] [Z] [H] [P] [N] [W] [P] [L] [C]
 1   2   3   4   5   6   7   8   9
*/

type stack struct {
	crates []string // bottom to top
}

func (s *stack) push(e string) {
	s.crates = append(s.crates, e)
}

func (s *stack) pushSlice(ee []string) {
	s.crates = append(s.crates, ee...)
}

func (s *stack) pop() string {
	e := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return e
}

func (s *stack) popSlice(size int) []string {
	ee := s.crates[len(s.crates)-size:]
	s.crates = s.crates[:len(s.crates)-size]
	return ee
}

func (s *stack) getTopElement() string {
	return s.crates[len(s.crates)-1]
}

func newStack() *stack {
	return &stack{
		crates: []string{},
	}
}

type stacksCollection struct {
	stacks map[int]*stack
	re     *regexp.Regexp
}

func newStacksCollection() *stacksCollection {
	re, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}

	m := make(map[int]*stack)
	for i := 1; i < 10; i++ {
		m[i] = newStack()
	}
	inits := [][]string{
		{"M", "J", "C", "B", "F", "R", "L", "H"},
		{"Z", "C", "D"},
		{"H", "J", "F", "C", "N", "G", "W"},
		{"P", "J", "D", "M", "T", "S", "B"},
		{"N", "C", "D", "R", "J"},
		{"W", "L", "D", "Q", "P", "J", "G", "Z"},
		{"P", "Z", "T", "F", "R", "H"},
		{"L", "V", "M", "G"},
		{"C", "B", "G", "P", "F", "Q", "R", "J"}}

	sc := &stacksCollection{stacks: m, re: re}
	sc.loadInitialStacks(inits)

	return sc
}

func (sc *stacksCollection) loadInitialStacks(cratesLists [][]string) {
	for i, list := range cratesLists {
		for _, e := range list {
			sc.stacks[i+1].push(e)
		}
	}
}

func (sc *stacksCollection) parseMoves(s string) (int, int, int) {

	digits := sc.re.FindAllString(s, -1)

	repeats, err := strconv.Atoi(digits[0])
	if err != nil {
		panic(err)
	}
	from, err := strconv.Atoi(digits[1])
	if err != nil {
		panic(err)
	}
	to, err := strconv.Atoi(digits[2])
	if err != nil {
		panic(err)
	}
	return repeats, from, to
}

func (sc *stacksCollection) performMoves(s string) {
	repeats, from, to := sc.parseMoves(s)

	for i := 0; i < repeats; i++ {
		e := sc.stacks[from].pop()
		sc.stacks[to].push(e)
	}
}

func (sc *stacksCollection) performSecondPartMoves(s string) {
	count, from, to := sc.parseMoves(s)

	ee := sc.stacks[from].popSlice(count)
	sc.stacks[to].pushSlice(ee)
}

func (sc *stacksCollection) printTopCrates() {
	for i := 1; i < 10; i++ {
		fmt.Print(sc.stacks[i].getTopElement())
	}
	fmt.Print("\n")
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := newStacksCollection()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sc.performMoves(scanner.Text())
	}

	fmt.Print("Final top crates (part 1): ")
	sc.printTopCrates()

	f.Seek(0, io.SeekStart)

	sc = newStacksCollection()
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		sc.performSecondPartMoves(scanner.Text())
	}

	fmt.Print("Final top crates (part 2): ")
	sc.printTopCrates()
}
