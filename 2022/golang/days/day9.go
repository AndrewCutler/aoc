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
					tail.y = head.y
					if head.x-tail.x > 1 {
						tail.x++
						visited_points[tail] = true
					}
				}
			}
		case "U":
			{
				fmt.Printf("Up: %v\n", distance)
				for ; distance > 0; distance-- {
					head.y++
					tail.x = head.x
					if head.y-tail.y > 1 {
						tail.y++
						visited_points[tail] = true
					}
				}
			}
		case "L":
			{
				for ; distance > 0; distance-- {
					head.x--
					tail.y = head.y
					if head.x-tail.x < 1 {
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
					tail.x = head.x
					if head.y-tail.y < 1 {
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

	cols := 0
	rows := 0
	for key := range visited_points {
		if rows < key.x {
			rows = key.x + 1
		}
		if cols < key.y {
			cols = key.y + 1
		}
	}

	for i := rows; i > 0; i-- {
		line := ""
		for j := cols; j > 0; j-- {
			curr := Point{x: i, y: j}
			if visited_points[curr] {
				line += "#"
			} else {
				line += "*"
			}
		}
		fmt.Println(line)
	}

	fmt.Println(len(visited_points))
}
