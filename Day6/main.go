package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(file string) string {
	input, err := os.Open(file)
	if err != nil {
		panic("there was an issue opening that file")
	}

	scanner := bufio.NewScanner(input)

	var data string

	for scanner.Scan() {
		data = scanner.Text()
	}

	return data
}

func partOne(dataStream string) {
	// we will use two pointers to create a sliding window to check for duplicate chars
	// stop the window when we have a window with no duplicates
	// return starting window index + 1

	type exists struct{}

	// loop through string
	for i := 0; i < len(dataStream)-3; i++ {
		ans := i + 4
		d := dataStream
		checkChars := map[string]exists{}
		window := string(d[i]) + string(d[i+1]) + string(d[i+2]) + string(d[i+3])

		for _, s := range window {
			// fmt.Println(window)
			if _, ok := checkChars[string(s)]; ok {
				break
			} else {
				checkChars[string(s)] = exists{}
			}
		}
		if len(checkChars) == 4 {
			fmt.Println(ans)
			return
		}
	}
}

func partTwo(dataStream string) {

	// same concept as part one but with a window of 14 chars instead of four
	type exists struct{}

	// loop through string
	for i := 0; i < len(dataStream)-13; i++ {
		ans := i + 14
		d := dataStream
		checkChars := map[string]exists{}

		//There HAS to be a better way to do this
		window := string(d[i]) + string(d[i+1]) + string(d[i+2]) + string(d[i+3]) + string(d[i+4]) + string(d[i+5]) + string(d[i+6]) + string(d[i+7]) + string(d[i+8]) + string(d[i+9]) + string(d[i+10]) + string(d[i+11]) + string(d[i+12]) + string(d[i+13])

		for _, s := range window {
			// fmt.Println(window)
			if _, ok := checkChars[string(s)]; ok {
				break
			} else {
				checkChars[string(s)] = exists{}
			}
		}
		if len(checkChars) == 14 {
			fmt.Println(ans)
			return
		}
	}
}

func main() {
	data := readFile("input.txt")
	partOne(data)
	partTwo(data)
}
