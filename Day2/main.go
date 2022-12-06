package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// C Y
	// B Y
	// C Y
	// C Y
	// B Y

	type Outcomes struct {
		t, w, l string
		val     int
	}
	rock := map[string]int{
		"X": 3,
		"Y": 6,
		"Z": 0,
	}

	paper := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	scissors := map[string]int{
		"X": 6,
		"Y": 0,
		"Z": 3,
	}

	scores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	outcomes := map[string]Outcomes{
		"X": {"A", "C", "B", 1},
		"Y": {"B", "A", "C", 2},
		"Z": {"C", "B", "A", 3},
	}

	// partTwoScores := map[string]int{
	// 	"X" : 0,
	// 	"Y" : 3,
	// 	"Z" : 6,
	// 	"rock" : 1,
	// 	"paper" : 2,
	// 	"scissors" : 3,
	// }

	// winLossDraw := map[string]string {
	// 	"rockW" : "paper",
	// 	"rockL" : "scissors",
	// 	"rockT" : "rock",
	// 	"paperW" : "scissors",
	// 	"paperL" : "rock",
	// 	"paperT" : "paper",
	// 	"scissorsW" : "rock",
	// 	"scissorsL" : "paper",
	// 	"scissorsT" : "scissors",
	// 	"A" : "rock",
	// 	"B" : "paper",
	// 	"C" : "scissors",
	// }

	input, err := os.Open("input.txt")
	if err != nil {
		panic("cannot open file")
	}

	// fmt.Println("this is input before scanner: ", input)

	scanner := bufio.NewScanner(input)

	var lines []string
	partOneTotal := 0
	partTwoTotal := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// lines = append(lines, " ")

	// fmt.Println(lines)

	for _, line := range lines {
		choices := strings.Split(line, " ")

		c := choices[1]
		z := choices[0]
		//part one
		switch z {
		case "A":
			partOneTotal += scores[c] + rock[c]
		case "B":
			partOneTotal += scores[c] + paper[c]
		case "C":
			partOneTotal += scores[c] + scissors[c]
		}

		//part two
		switch c {
		case "X":
			for _, r := range outcomes {
				if z == r.l {
					partTwoTotal += r.val
				}
			}
		case "Y":
			for _, r := range outcomes {
				if z == r.t {
					partTwoTotal += (r.val + 3)
				}
			}
		case "Z":
			for _, r := range outcomes {
				if z == r.w {
					partTwoTotal += (r.val + 6)
				}
			}
		}

	}

	fmt.Println(partOneTotal)
	fmt.Println(partTwoTotal)
	input.Close()
}
