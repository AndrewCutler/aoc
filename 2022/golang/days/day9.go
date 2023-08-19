package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func DayNine() {
	file, err := os.Open("../data/day9.test.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	curr := Point{x: 0, y: 0}
	// visited_points := []Point{Point{x: 0, y: 0}}

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")
		direction := args[0]
		distance, err := strconv.Atoi(args[1])

		if err != nil {
			panic(err)
		}

		switch direction {
		case "R":
			{
				fmt.Printf("Right: %v\n", distance)
				for ; distance > 0; distance-- {
					curr.x++
				}
			}
		case "U":
			{
				fmt.Printf("Up: %v\n", distance)
			}
		case "L":
			{
				fmt.Printf("Left: %v\n", distance)
			}
		case "D":
			{
				fmt.Printf("Down: %v\n", distance)
			}
		default:
			{
				panic("Invalid direction")
			}
		}
	}
}
