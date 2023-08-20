package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Find signal strength at cycles 20, 60, 100, 140, 180 and 220.
// Signal strength = cycle # * current value in register
func DayTen() {
	file, err := os.Open("../data/day10.test.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	register_value := 0
	for cycle := 0; scanner.Scan(); cycle++ {
		command := scanner.Text()

		if command == "noop" {
			continue
		}

		if !strings.HasPrefix(command, "addx") {
			err := fmt.Sprintf("Invalid command: %v\n", command)
			panic(err)
		}

		value, err := strconv.Atoi(strings.Split(command, " ")[1])

		if err != nil {
			panic(err)
		}

		fmt.Println(value)
		cycle++
		register_value += value
	}
}
