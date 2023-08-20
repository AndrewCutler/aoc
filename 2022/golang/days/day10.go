package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var target_cycles = []int{20, 60, 100, 140, 180, 220}

func is_target_cycle(cycle int) bool {
	for _, curr := range target_cycles {
		if cycle == curr {
			return true
		}
	}

	return false
}

func check_target_cycles(cycle int, register_value int, sum *int) {
	if is_target_cycle(cycle) {
		signal_strength := cycle * register_value
		*sum += signal_strength
	}
}

func draw(register_value int, x_position *int, screen *string) {
	if register_value == *x_position || register_value-1 == *x_position || register_value+1 == *x_position {
		*screen += "# "
	} else {
		*screen += ". "
	}

	*x_position++
	if *x_position == 40 {
		*screen += "\n"
		*x_position = 0
	}
}

func DayTen() {
	file, err := os.Open("../data/day10.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	register_value := 1
	signal_strengths_sum := 0
	screen := ""
	x_position := 0
	
	for cycle := 1; scanner.Scan(); cycle++ {
		command := scanner.Text()

		draw(register_value, &x_position, &screen)
		check_target_cycles(cycle, register_value, &signal_strengths_sum)

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

		draw(register_value, &x_position, &screen)
		check_target_cycles(cycle, register_value, &signal_strengths_sum)

		register_value += value
	}

	fmt.Printf("Part one: %v\n", signal_strengths_sum)
	fmt.Printf("Part two: \n%v\n", screen)
}
