package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	day1()
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