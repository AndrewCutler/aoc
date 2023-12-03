package main

import (
	days "aoc/days"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("Invalid arguments. Day must be specified.")
	}

	switch day := args[1]; day {
	case "1":
		{
			days.DayOne()
		}
	default:
		panic("Invalid day specified. Valid options are: 1, 5, 6.1, 6.2, 7.")
	}
}
