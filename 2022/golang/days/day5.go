package days

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// if I continue to learn Go, I'm sure this will make me shudder in 6 months
func DayFive() {
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
}
