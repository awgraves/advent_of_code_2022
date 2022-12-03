package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var counts []int

	c := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			counts = append(counts, c)
			c = 0
			continue
		}
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		c += i
	}

	max := 0
	for _, i := range counts {
		if i > max {
			max = i
		}
	}
	fmt.Printf("Most calories: %d\n", max) // part 1 solution

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	topThree := 0
	for _, i := range counts[:3] {
		topThree += i
	}
	fmt.Printf("Top three combined: %d\n", topThree) // part 2 solution
}
