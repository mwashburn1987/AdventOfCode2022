package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readFile("input.txt")
	tree := createTree(data)
	fmt.Println("Part one answer is: ", partOne(tree))

	spaceNeeded := tree.size - 40_000_000
	minFile := math.MaxInt

	fmt.Println("Part two answer is: ", partTwo(tree, spaceNeeded, minFile))
}

// node properties for our tree
type Node struct {
	name      string
	size      int
	contains  bool
	parentDir *Node
	children  []*Node
}

// methods for nodes

// adds a file node
func (n *Node) addChild(name string, size int) {
	n.children = append(n.children, &Node{name: name, size: size, parentDir: n})
	n.addSize(size)
}

// adds a directory node
func (n *Node) addDirectory(name string) {
	n.children = append(n.children, &Node{name: name, parentDir: n, children: []*Node{}, contains: true})
}

// this will recursively add size up our file path to parent directories
func (n *Node) addSize(size int) {
	n.size += size
	if n.parentDir != nil {
		n.parentDir.addSize(size)
	}
}

func readFile(filename string) []string {
	f, _ := os.Open(filename)
	defer f.Close()

	sc := bufio.NewScanner(f)
	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

func createTree(lines []string) *Node {

	// create the root of our tree as we know it starts with "/"
	root := &Node{name: "/", size: 0, children: []*Node{}, contains: true}

	// set the current node we're operating on as root node
	cur := root

	// loop through lines and conditionally operate

	for i := 0; i < len(lines); i++ {

		line := lines[i]

		if line == "$ ls" {
			for {
				// check next line to see if it's a command, if so exit early
				if i == (len(lines)-1) || lines[i+1][0] == '$' {
					break
				}
				// now we can move to next line
				i++

				line = lines[i]
				contents := strings.Split(line, " ")

				if contents[0] == "dir" {
					cur.addDirectory(contents[1])
				} else {
					size, err := strconv.Atoi(contents[0])
					if err != nil {
						panic(err)
					}
					cur.addChild(contents[1], size)
				}
			}
		} else if line == "$ cd .." {
			cur = cur.parentDir

		} else {
			// all that's left is to check for more directories
			content := strings.Split(line, " ")

			for _, node := range cur.children {
				if node.name == content[2] {
					cur = node
				}
			}
		}

	}
	return root
}

func partOne(n *Node) int {
	total := 0

	if n.size <= 100000 {
		total += n.size
	}

	// recursively work through directories
	for _, node := range n.children {
		if node.contains {
			total += partOne(node)
		}
	}
	return total
}

func partTwo(n *Node, spaceNeeded int, minFile int) int {
	// search tree for directories, if file size is >=  space needed, check if it's the smallest file that
	// accomplishes our goal, if so update min and then finally return min

	curMin := minFile

	if n.contains {
		if n.size >= spaceNeeded {
			fmt.Println(n.name, n.size)
			curMin = int(math.Min(float64(curMin), float64(n.size)))
			fmt.Println("current minFile: ", curMin)
		} else {
			return curMin
		}
	}
	for _, node := range n.children {
		curMin = partTwo(node, spaceNeeded, curMin)
	}

	return int(curMin)
}
