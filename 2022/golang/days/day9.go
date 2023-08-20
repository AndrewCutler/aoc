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
	file, err := os.Open("../data/day9.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	head := Point{x: 0, y: 0}
	tail := Point{x: 0, y: 0}
	visited_points := map[Point]bool{tail: true}

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
				for ; distance > 0; distance-- {
					head.x++
					if head.x-tail.x > 1 {
						tail.y = head.y
						tail.x++
						visited_points[tail] = true
					}
				}
			}
		case "U":
			{
				for ; distance > 0; distance-- {
					head.y++
					if head.y-tail.y > 1 {
						tail.x = head.x
						tail.y++
						visited_points[tail] = true
					}
				}
			}
		case "L":
			{
				for ; distance > 0; distance-- {
					head.x--
					if tail.x-head.x > 1 {
						tail.y = head.y
						tail.x--
						visited_points[tail] = true
					}
				}
			}
		case "D":
			{
				fmt.Printf("Up: %v\n", distance)
				for ; distance > 0; distance-- {
					head.y--
					if tail.y-head.y > 1 {
						tail.x = head.x
						tail.y--
						visited_points[tail] = true
					}
				}
			}
		default:
			{
				panic("Invalid direction")
			}
		}
	}

	fmt.Println(len(visited_points))
}
