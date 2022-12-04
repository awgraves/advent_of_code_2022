package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getRanges(line string) [2][2]int {
	intRanges := [2][2]int{}

	strRanges := strings.Split(line, ",")
	for i := 0; i < 2; i++ {
		r := strRanges[i]
		strInts := strings.Split(r, "-")
		for j, s := range strInts {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			intRanges[i][j] = num
		}
	}
	return intRanges
}

func isContainment(rs [2][2]int) bool {
	// containment def: one range completely encompasses another
	pairOneContains := (rs[0][0] <= rs[1][0]) && (rs[0][1] >= rs[1][1])
	pairTwoContains := (rs[0][0] >= rs[1][0]) && (rs[0][1] <= rs[1][1])
	return pairOneContains || pairTwoContains
}

func isOverlap(rs [2][2]int) bool {
	// overlap def: any number in one range overlaps with the other range

	noOverlap := (rs[0][0] < rs[1][0] && rs[0][1] < rs[1][0]) || (rs[1][0] < rs[0][0] && rs[1][1] < rs[0][0])

	return !noOverlap
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	containmentCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ranges := getRanges(scanner.Text())
		if o := isContainment(ranges); o {
			containmentCount++
		}
	}
	fmt.Printf("Containment count: %d\n", containmentCount)

	file.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(file)
	overlapCount := 0
	for scanner.Scan() {
		ranges := getRanges(scanner.Text())
		if o := isOverlap(ranges); o {
			overlapCount++
		}
	}
	fmt.Printf("Overlap count: %d\n", overlapCount)
}
