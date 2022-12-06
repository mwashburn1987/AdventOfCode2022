package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := readInput("input.txt")
	partOne(input)
	partTwo(input)
}

func readInput(file string) []string {
	content, err := os.Open("input.txt")

	if err != nil {
		panic("could not open file")

	}

	scanner := bufio.NewScanner(content)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lines = append(lines, " ")
	return lines
}

func partOne(lines []string) {
	//scan each element, convert to float, add together and check if that's current max
	c := 0
	m := 0

	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		c += n
		if line == "" {
			m = int(math.Max(float64(m), float64(c)))
			c = 0
		}
	}
	fmt.Println(m)
}

func partTwo(lines []string) {
	total := 0
	curr := 0
	highestVals := [4]int{0, 0, 0, 0}

	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		curr += n
		if line == "" {
			highestVals[3] = curr
			sort.Sort(sort.Reverse((sort.IntSlice(highestVals[:]))))
			curr = 0
		}
	}
	total = highestVals[0] + highestVals[1] + highestVals[2]

	fmt.Println(total)
}
