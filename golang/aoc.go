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
	day5()
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
			fmt.Println("done reading file")
			break
		}
	}

	// now we have all the moves so let's move them
	for i := 0; i < len(moves); i++ {
		charsToMove := make([]string, 0)
		source := moves[i].source - 1
		for ; moves[i].count > 0; moves[i].count-- {
			if len(crates[source]) > 0 {
				charsToMove = append(charsToMove, crates[source][0])
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
