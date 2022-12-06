package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func to get the letter value of common char
func letterVal(l string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(alphabet, l) + 1
}

func readFile(file string) []string {
	input, err := os.Open(file)
	if err != nil {
		panic("could not open given file")
	}

	scanner := bufio.NewScanner(input)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		// lines = scanner.Text()
	}

	return lines
}

func solver(data []string) {
	//take the line and split in two
	//create a map to act as a set and find the common character
	//find value of common characters priority and add it to total
	//return total

	total := 0
	//empty struct for map to consume less memory
	type exists struct{}
	badgeSet := map[rune]exists{}
	badgeCommon := map[rune]exists{}
	threeTotal := 0
	lineCounter := 1

	for _, line := range data {
		set := map[rune]exists{}

		var common string

		half := (len(line) / 2) - 1
		for i, char := range line {
			if i <= half {
				set[char] = exists{}
				// badgeSet[char] = exists{}
			} else {
				if _, ok := set[char]; ok {
					common = string(char)
					break
				}
				// badgeSet[char] = exists{}
			}

		}
		total += letterVal(common)
	}

	//part two
	for _, line := range data {
		if lineCounter == 1 {
			//second set to check for common char in all 3 lines

			for _, char := range line {
				badgeSet[char] = exists{}
			}
			lineCounter++
		} else if lineCounter == 2 {
			for _, char := range line {
				if _, ok := badgeSet[char]; ok {
					badgeCommon[char] = exists{}
				}
			}
			lineCounter++
		} else if lineCounter == 3 {
			for _, char := range line {
				if _, ok := badgeCommon[char]; ok {
					threeTotal += letterVal(string(char))
					badgeCommon = map[rune]exists{}
					badgeSet = map[rune]exists{}
					lineCounter = 1
					break
				}
			}
		}
	}
	fmt.Println(total)
	fmt.Println(threeTotal)

	// total := 0
	// totalGroups := 0

	// groupsOfThree := []string{}

	// for _, line := range data {
	// 	half := (len(line) / 2)
	// 	sub := []string{line[:half], line[half:]}

	// 	for _, r := range sub[0] {
	// 		if strings.Contains(sub[1], string(r)) {
	// 			total += letterVal(string(r))
	// 			break
	// 		}
	// 	}
	// 	groupsOfThree = append(groupsOfThree, line)

	// 	if len(groupsOfThree) == 3 {
	// 		for _, r := range groupsOfThree[0] {
	// 			if strings.Contains(groupsOfThree[1], string(r)) && strings.Contains(groupsOfThree[2], string(r)) {
	// 				totalGroups += letterVal(string(r))
	// 				groupsOfThree = []string{}
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(total)
	// fmt.Println(totalGroups)
}

func main() {
	data := readFile("input.txt")
	solver(data)

}
