package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func getCommonRune(line string) rune {
	mid := len(line) / 2

	c1 := line[:mid]
	c2 := line[mid:]

	seen := make(map[rune]bool)
	for _, r := range c1 {
		seen[r] = true
	}

	for _, r := range c2 {
		_, ok := seen[r]
		if ok {
			return r
		}
	}
	panic("Not found")
}

func getBadgeRune(threeLines []string) rune {
	seen := make(map[rune]int)

	for _, r := range threeLines[0] {
		seen[r] = 1
	}
	for _, r := range threeLines[1] {
		_, ok := seen[r]
		if !ok {
			continue
		}
		seen[r] = 2
	}
	for _, r := range threeLines[2] {
		i, ok := seen[r]
		if !ok {
			continue
		}
		if i == 2 {
			return r
		}
	}
	panic("not found")
}

func getPriorityOfRune(r rune) int {
	if string(r) == strings.ToLower(string(r)) {
		return int(r - 'a' + 1)
	}
	return int(r - 'A' + 27)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPriority := 0
	for scanner.Scan() {
		r := getCommonRune(scanner.Text())
		totalPriority += getPriorityOfRune(r)
	}
	fmt.Printf("Total priority: %d\n", totalPriority)

	// part 2
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(file)
	totalPriority = 0
Mainloop:
	for {
		group := []string{}
		for i := 0; i < 3; i++ {
			if c := scanner.Scan(); !c {
				break Mainloop
			}
			group = append(group, scanner.Text())
		}
		r := getBadgeRune(group)
		totalPriority += getPriorityOfRune(r)
	}
	fmt.Printf("Total badge priority: %d\n", totalPriority)
}
