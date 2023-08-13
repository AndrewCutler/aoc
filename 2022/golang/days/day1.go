package day_one

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func DayOne() {
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
	part_one := sums[len-1]
	second := sums[len-2]
	third := sums[len-3]
	part_two := part_one + second + third


	fmt.Println("Day one:")
	fmt.Printf("\tPart one: %v\n\tPart two: %v\n", part_one, part_two)
}
