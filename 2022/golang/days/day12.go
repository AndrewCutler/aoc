package days

import (
	"bufio"
	"fmt"
	"os"
)

// recursive function that accepts a square (top, right, bottom or left).
// if square is nil or square.height - current.height > 1, return.
// if square == Z, we've found a successful route.

// var paths []Path

type TopoSquare struct {
	steps  []int // each step is an int; 0 for up, 1 for right, 2 for down, 3 for left
	height int
	symbol string // a-z, S, E
	top    *Square
	right  *Square
	bottom *Square
	left   *Square
}

func visit(curr *TopoSquare) {
	if curr == nil {
		return
	}
}

func DayTwelve() {
	fmt.Println("day 12")
	file, err := os.Open("../data/day12.test.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for _, c := range line {

			chars := []rune("abcdefghijklmnopqrstuvwxyzSE")
			switch c {
			case chars[0]:
				{
					fmt.Println(chars[0], "a")
				}
			}
			// fmt.Println(c)
		}
	}
}
