package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func DayOne() {
	file, err := os.Open("../data/day1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		var first rune
		var second rune
		for _, char := range line {
			if unicode.IsDigit(char) {
				if !unicode.IsPrint(first) {
					first = char
				}
				second = char
			}
		}

		combined := string([]rune{first, second})
		converted, err := strconv.Atoi(combined)
		if err != nil {
			panic(err)
		}

		total += converted
	}

	fmt.Printf("Day one: %d\n", total)
}
