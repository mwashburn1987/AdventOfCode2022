package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

// check to see if stack contains any values
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// push a value to top of stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// pop a value off the stack and return that value. Return false if stack is empty
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	ind := len(*s) - 1 // gets the last (top) index of the slice
	ele := (*s)[ind]   // assigns the last element to a variable
	*s = (*s)[:ind]    // reassign the stack missing last element
	return ele, true
}

// read file and return usable data
func readFile(file string) ([]string, [][]int) {
	input, err := os.Open(file)
	if err != nil {
		panic("unable to open selected file")
	}

	scanner := bufio.NewScanner(input)

	var instructions [][]int
	var a, b, c int

	lines := []string{}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		lines = append(lines, scanner.Text())
	}

	for scanner.Scan() {
		_, err = fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &a, &b, &c)
		if err != nil {
			fmt.Println("there was an error scanning file: ", err)
		}
		nums := []int{a, b, c}
		instructions = append(instructions, nums)
	}
	return lines, instructions

}

// takes in data read in from readFile and generates a solution
func solverPartOne(lines []string, instructions [][]int) {

	stacksString := strings.Split(lines[len(lines)-1], " ")

	numOfStacks, err := strconv.Atoi(stacksString[len(stacksString)-2])
	if err != nil {
		fmt.Println("there was an error finding number of stacks: ", err)
	}

	stacks := []Stack{}

	//generate stacks
	for i := 1; i <= numOfStacks; i++ {
		stacks = append(stacks, Stack{})
	}
	//get rid of last line for data manipulation
	lines = lines[:len(lines)-1]

	//loop through stack data lines and generate stacks
	for i := len(lines) - 1; i >= 0; i-- {
		//index to add to proper stack slice
		j := 0
		for k := 1; k <= len(lines[i]); k += 4 {
			if string(lines[i][k]) == " " {
				j++
				continue
			}
			stacks[j] = append(stacks[j], string(lines[i][k]))
			j++
		}
	}

	for _, ins := range instructions {
		//get our to and from sstack out of instructions, subtract one to match indexes
		cratesToMove := ins[0]
		from := stacks[ins[1]-1]
		to := stacks[ins[2]-1]

		stacks[ins[1]-1], stacks[ins[2]-1] = moveCrates(from, to, cratesToMove)
	}

	results := ""

	for _, s := range stacks {
		results += s[len(s)-1]
	}

	fmt.Println(results)

}

func solverPartTwo(lines []string, instructions [][]int) {

	stacksString := strings.Split(lines[len(lines)-1], " ")

	numOfStacks, err := strconv.Atoi(stacksString[len(stacksString)-2])
	if err != nil {
		fmt.Println("there was an error finding number of stacks: ", err)
	}

	stacks := []Stack{}

	//generate stacks
	for i := 1; i <= numOfStacks; i++ {
		stacks = append(stacks, Stack{})
	}

	//get rid of last line for data manipulation
	lines = lines[:len(lines)-1]

	//loop through stack data lines and generate stacks
	for i := len(lines) - 1; i >= 0; i-- {
		//index to add to proper stack slice
		j := 0
		for k := 1; k <= len(lines[i]); k += 4 {
			if string(lines[i][k]) == " " {
				j++
				continue
			}
			stacks[j] = append(stacks[j], string(lines[i][k]))
			j++
		}
	}

	for _, ins := range instructions {
		n := ins[0]
		stacks[ins[1]-1], stacks[ins[2]-1] = moveCratesPartTwo(stacks[ins[1]-1], stacks[ins[2]-1], n)

	}

	results := ""

	for _, s := range stacks {
		results += s[len(s)-1]
	}

	fmt.Println(results)
}
func moveCrates(from Stack, to Stack, numOfCrates int) (Stack, Stack) {
	// fmt.Sprintf("the input crates are - From: %v - To: %v", from, to)
	for i := 1; i <= numOfCrates; i++ {
		if crate, ok := from.Pop(); ok {
			to.Push(crate)
		}
	}
	return from, to
}

// to solve part two with multiple crates picked up at a time
func moveCratesPartTwo(from Stack, to Stack, numOfCrates int) (Stack, Stack) {
	// fmt.Println("from: ", from, "to: ", to, "crateNum: ", numOfCrates)
	tmp := []string{}
	// this will move from top of stack down
	for i := 1; i <= numOfCrates; i++ {
		c, _ := from.Pop()
		tmp = append(tmp, c)
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		to.Push(tmp[i])
	}
	// fmt.Println("from after pop: ", from, "to after push: ", to)
	return from, to
}

func main() {
	lines, instructions := readFile("input.txt")
	solverPartOne(lines, instructions)
	solverPartTwo(lines, instructions)

}
