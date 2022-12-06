package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(file string) []string {
	input, err := os.Open(file)
	if err != nil {
		panic("could not open given file")
	}

	scanner := bufio.NewScanner(input)

	lines := []string{}
	for scanner.Scan() {
		var a, b, c, d int
		_, errr := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &a, &b, &c, &d)
		fmt.Println(a, b, c, d)
		if errr != nil {
			panic("error parsing lines with Sscanf")
		}
		lines = append(lines, scanner.Text())
		// lines = scanner.Text()
	}

	fmt.Println("hello world")
	return lines
}
func solver(data []string) {
	totalPairs := 0
	for _, line := range data {
		groups := strings.Split(line, ",")

		firstRange := strings.Split(groups[0], "-")
		firstOne, _ := strconv.Atoi(firstRange[0])
		firstTwo, _ := strconv.Atoi(firstRange[1])

		secondRange := strings.Split(groups[1], "-")
		secondOne, _ := strconv.Atoi(secondRange[0])
		secondTwo, _ := strconv.Atoi(secondRange[1])

		intRanges := []int{firstOne, firstTwo, secondOne, secondTwo}
		// fmt.Println(intRanges)

		if (intRanges[0] == intRanges[2]) || (intRanges[1] == intRanges[3]) {
			// fmt.Println("first int: ", intRanges[0], "second int: ", intRanges[1])
			totalPairs++
			continue
		}
		// beginningRangeMin := math.Min(float64(intRanges[0]), float64(intRanges[1]))

		if intRanges[0] <= intRanges[2] {

			if intRanges[1] >= intRanges[3] {
				// fmt.Println("element 1: ", intRanges[1], "element 3: ", intRanges[3])
				totalPairs++
				continue
				//comment out below else if for part one answer only
			} else if intRanges[2] >= intRanges[0] && intRanges[2] <= intRanges[1] {
				totalPairs++
				continue
			} else {
				continue
			}
		} else if intRanges[2] < intRanges[0] {
			if intRanges[3] >= intRanges[1] {
				// fmt.Println("element 3: ", intRanges[3], "element 1: ", intRanges[1])
				totalPairs++
				continue
				//comment out below else if for part one answer only
			} else if intRanges[0] >= intRanges[2] && intRanges[0] <= intRanges[3] {
				totalPairs++
				continue
			} else {
				continue
			}
		}
	}
	fmt.Println(totalPairs)

}
func main() {
	data := readFile("input.txt")

	// data := readFile("sampleData.txt")
	solver(data)
}
