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
	type move struct {
		count  int
		source int
		dest   int
	}
	var stacks [][]string

	file, err := os.Open("../data/day5.test.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var lineBytes []byte

	var numOfCols float64
	// put stack in proper columns
	for i := 1; ; i++ {
		lineBytes, err = reader.ReadBytes('\n')
		length := len(lineBytes)

		if err != nil {
			fmt.Println(err)
			break
		}

		// get number of columns, but only do so once
		if i == 1 {
			numOfCols = math.Ceil(float64(length / 4))
			stacks = make([][]string, int(numOfCols))
		}
		// fmt.Printf("Number of columns: %v\n", numOfCols)

		line := string(lineBytes[:])
		// do until first line without [
		if strings.Contains(line, "[") == false {
			break
		}

		offset := 1
		step := 4
		for j := 0; j < int(numOfCols); j++ {
			currentPosition := offset + (step * j)
			currentByte := lineBytes[currentPosition]
			prev := stacks[j]
			currentChar := string(currentByte)
			if currentChar != " " {
				stacks[j] = append(prev, string(currentByte))
			}
		}
	}
	// here we have each column in a string array, excluding whitespace
	// fmt.Println(stacks)

	if err != io.EOF && err != nil {
		fmt.Println(err)
	}

	// do moving of stacks
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
		// fmt.Printf("i = %v; %v\n", i, line)
		if strings.HasPrefix(line, "move") {
			// fmt.Printf("i = %v; %v\n", i, line)

			// get numbers: first is count, second is source, third is destination
			re := regexp.MustCompile("\\d")
			nums := re.FindAllString(line, 3)
			count, countErr := strconv.Atoi(nums[0])
			source, sourceErr := strconv.Atoi(nums[1])
			dest, destErr := strconv.Atoi(nums[2])
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

	// now we have all the moves
	fmt.Println(stacks)
	for i := 0; i < len(moves); i++ {
		charsToMove := make([]string, 0)
		source := moves[i].source-1
		// fmt.Println(stacks[source])
		for j := 0; moves[i].count > 0; moves[i].count-- {
			// fmt.Printf("current char: %v\n", stacks[source][j])
			charsToMove = append(charsToMove, stacks[source][j])
			// pop from stack here
			j++ // how do I do this in the for post statement?
		}
		fmt.Println(charsToMove)
		// have charsToMove for current move; now pop from stack and push to dest
	}
	fmt.Println(stacks)
}
