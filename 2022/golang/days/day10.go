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
	target_cycles := []int{20, 60, 100, 140,180,220}
	file, err := os.Open("../data/day10.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	register_value := 1
	signal_strengths_sum := 0
	for cycle := 1; scanner.Scan(); cycle++ {
		command := scanner.Text()

		for _, c := range target_cycles {
			if cycle == c {
				signal_strength := cycle * register_value
				signal_strengths_sum += signal_strength
			}
		}

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

		cycle++
		for _, c := range target_cycles {
			if cycle == c {
				signal_strength := cycle * register_value
				signal_strengths_sum += signal_strength
			}
		}

		register_value += value
	}

	fmt.Printf("Part one: %v\n", signal_strengths_sum)
}
