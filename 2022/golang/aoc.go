package main

import (
	days "aoc/days"
	"math/big"
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
	case "5":
		{
			days.DayFive()
		}
	case "6.1":
		{
			days.DaySix(4)
		}
	case "6.2":
		{
			days.DaySix(14)
		}
	case "7":
		{
			days.DaySeven()
		}
	case "8":
		{
			days.DayEight()
		}
	case "9":
		{
			days.DayNine()
		}
	case "10":
		{
			days.DayTen()
		}
	case "11.1":
		{
			days.DayEleven(big.NewInt(3), 20)
		}
	case "11.2":
		{
			days.DayEleven(big.NewInt(3), 10000)
		}
	case "12.1":
		{
			days.DayTwelve()
		}
	default:
		panic("Invalid day specified. Valid options are: 1, 5, 6.1, 6.2, 7.")
	}
}
