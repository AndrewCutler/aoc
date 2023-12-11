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
	case "2":
		{
			days.DayTwo()
		}
	default:
		panic("Invalid day specified. Valid options are: 1, 2.")
	}
}
