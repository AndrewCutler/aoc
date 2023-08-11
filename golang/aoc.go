package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// day1()
	// day5()
	// day6(4)  // part one
	// day6(14) // part two
	day7()
}

func day1() {
	file, err := os.Open("../data/day1.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	curr := 0
	sums := make([]int, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			sums = append(sums, curr)
			curr = 0
			continue
		}

		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		curr += val
	}
	sort.Ints(sums)
	len := len(sums)
	partOne := sums[len-1]
	second := sums[len-2]
	third := sums[len-3]
	partTwo := partOne + second + third

	if partOne != 68775 {
		panic("part one is incorrect")
	}
	if partTwo != 202585 {
		panic("part two is incorrect")
	}

	fmt.Println("Huzzah!")
}

// if I continue to learn Go, I'm sure this will make me shudder in 6 months
func day5() {
	var crates [][]string

	file, err := os.Open("../data/day5.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var lineBytes []byte

	var numOfStacks float64
	// put crate in proper stacks
	for i := 1; ; i++ {
		lineBytes, err = reader.ReadBytes('\n')
		length := len(lineBytes)

		if err != nil {
			fmt.Println(err)
			break
		}

		// get number of stacks, but only do so once
		if i == 1 {
			numOfStacks = math.Ceil(float64(length / 4))
			crates = make([][]string, int(numOfStacks))
		}

		line := string(lineBytes[:])
		// do until first line without [
		if strings.Contains(line, "[") == false {
			break
		}

		offset := 1
		step := 4
		for j := 0; j < int(numOfStacks); j++ {
			currentPosition := offset + (step * j)
			currentByte := lineBytes[currentPosition]
			prev := crates[j]
			currentChar := string(currentByte)
			if currentChar != " " {
				crates[j] = append(prev, string(currentByte))
			}
		}
	}
	// here we have each column in a string array, excluding whitespace

	if err != io.EOF && err != nil {
		fmt.Println(err)
	}

	// figure out what goes where
	type move struct {
		count  int
		source int
		dest   int
	}
	i := 0
	moves := make([]move, 0)
	for {
		lineBytes, err = reader.ReadBytes('\n')

		if err != io.EOF && err != nil {
			fmt.Println(err)
			break
		}
		line := string(lineBytes[:])

		// for debugging
		if strings.HasPrefix(line, "move") {
			// get numbers: first is count, second is source, third is destination
			re := regexp.MustCompile("\\s\\d+\\s?")
			nums := re.FindAllString(line, 3)
			count, countErr := strconv.Atoi(strings.TrimSpace(nums[0]))
			source, sourceErr := strconv.Atoi(strings.TrimSpace(nums[1]))
			dest, destErr := strconv.Atoi(strings.TrimSpace(nums[2]))
			if countErr != nil || sourceErr != nil || destErr != nil {
				break
			}
			move := move{count: count, source: source, dest: dest}
			moves = append(moves, move)
		}
		i++
		if err == io.EOF {
			fmt.Print("Done reading file. ")
			break
		}
	}

	// now we have all the moves so let's move them
	fmt.Print("Moving crates...\n")
	for i := 0; i < len(moves); i++ {
		charsToMove := make([]string, 0)
		source := moves[i].source - 1
		for ; moves[i].count > 0; moves[i].count-- {
			if len(crates[source]) > 0 {
				// append for part one
				// charsToMove = append(charsToMove, crates[source][0])
				// prepend instead for part two
				charsToMove = append([]string{crates[source][0]}, charsToMove...)
				// pop from stack here
				crates[source] = crates[source][1:]
			}
		}
		// have charsToMove for current move; now push to dest
		dest := moves[i].dest - 1
		for len(charsToMove) > 0 {
			crates[dest] = append([]string{charsToMove[0]}, crates[dest]...)
			charsToMove = charsToMove[1:]
		}
	}

	// now we get the first element in each column
	topCrates := make([]string, 0)
	for i := 0; i < int(numOfStacks); i++ {
		if len(crates[i]) > 0 {
			topCrates = append(topCrates, crates[i][0])
		}
	}

	fmt.Printf("Part one: %v\n", topCrates)
	fmt.Println("Nice.")
}

func day6(packetsize int) {
	file, err := os.Open("../data/day6.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	marker := make([]byte, 0, packetsize)
	answer := 1
	for {
		char, err := reader.ReadByte()

		if err != nil {
			fmt.Println(err)
			break
		}

		// push to queue
		marker = append(marker, char)

		// once size is four, check for uniques
		// if success, break with answer
		// if failure, dequeue
		if len(marker) == packetsize {
			visited := make(map[byte]bool)
			for i, curr := range marker {
				if visited[curr] == true {
					// we have a failure
					marker = marker[1:]
					break
				} else if i == packetsize-1 {
					goto done
				} else {
					visited[curr] = true
				}
			}
		}
		answer++
	}

done:
	fmt.Printf("Solution: %v\n", answer)
}

type Node struct {
	size     int
	name     string
	visited  bool
	parent   *Node
	children []*Node
	// files File -- files here, as struct with name and size
}

func traverse(root *Node, sum *int) {
	if root == nil || root.name == "" || root.visited == true {
		return
	}
	root.visited = true
	if root.size <= 100000 {
		*sum += root.size
	}
	for _, child := range root.children {
		traverse(child, sum)
	}
	if root.parent.name != "" {
		traverse(root.parent, sum)
	}
}

func get_sums() {

}

func day7() {
	// if a directory's contents exceed 100,000, add to answer
	sum := 0

	file, err := os.Open("../data/day7.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var curr *Node
	i := 1
OUTER:
	for scanner.Scan() {
		i++

		line := strings.ReplaceAll(scanner.Text(), "\r\n", "")

		if strings.HasPrefix(line, "$ ls") || strings.HasPrefix(line, "dir") {
			continue
		}
		if strings.HasSuffix(line, "cd ..") {
			curr = curr.parent
			continue
		}
		if strings.HasPrefix(line, "$ cd") {
			name := strings.Split(line, " ")[2]

			if curr == nil {
				node := Node{
					name:     name,
					children: make([]*Node, 0),
					visited:  false,
				}
				curr = &node
				continue
			}

			temp := *curr
			for _, child := range curr.children {
				if child.name == name {
					child.parent = curr
					curr = child
					continue OUTER
				}
			}

			*curr = Node{
				name:     name,
				children: make([]*Node, 0),
				visited:  false,
			}

			curr.parent = &temp
			if temp.name != "" {
				temp.children = append(temp.children, curr)
			}
			continue
		}

		size, err := strconv.Atoi(strings.Split(line, " ")[0])

		if err != nil {
			panic(err)
		}
		curr.size += size
	}

	// return to root
	for {
		fmt.Println(curr.name)
		if curr.parent == nil {
			break
		}
		curr = curr.parent
	}

	// now sum all directory sizes <= 100000
	// traverse(curr, &sum)

	fmt.Printf("Part one: %v\n", sum)
}
