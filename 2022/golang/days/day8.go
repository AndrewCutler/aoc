package days

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Square struct {
	top          *Square
	right        *Square
	bottom       *Square
	left         *Square
	visited      bool
	height       int
	scenic_score int
	visible      bool
	id           string
}

// 2D array of row, col
// number of rows = number of lines per file
// number of cols = number of chars per line
// if row == 0, top = nil
// if row == (number of rows) - 1, bottom = nil
// if col == 0, left = nil
// if col == (number of chars per line) - 1, right = nil

// read every byte into 2D array

func DayEight() {
	file, err := os.Open("../data/day8.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	num_of_rows := 0
	var num_of_cols int
	for scanner.Scan() {
		line := scanner.Text()
		num_of_cols = len(line)
		num_of_rows++
	}

	grid := make([][]*Square, num_of_cols)
	for i := range grid {
		grid[i] = make([]*Square, num_of_rows)
	}

	file.Seek(0, io.SeekStart)

	scanner = bufio.NewScanner(file)

	zero_byte_offset := 48
	for line_num := 0; scanner.Scan(); line_num++ {
		line := scanner.Bytes()

		// first, allocate squares
		for i, b := range line {
			curr := Square{
				visited: false,
				height:  int(b) - zero_byte_offset,
				id:      string(line_num) + string(i),
			}

			grid[line_num][i] = &curr
		}
	}

	for row, _ := range grid {
		for col, _ := range grid[row] {
			curr := grid[row][col]
			if row != 0 {
				curr.top = grid[row-1][col]
			}
			if row != num_of_rows-1 {
				curr.bottom = grid[row+1][col]
			}
			if col != 0 {
				curr.left = grid[row][col-1]
			}
			if col != num_of_cols-1 {
				curr.right = grid[row][col+1]
			}
		}
	}

	for row_num := range grid {
		for _, square := range grid[row_num] {
			visible := false
			curr := square
			top_score := 0
			for {
				if curr.top == nil {
					visible = true
					break
				}
				if square.height <= curr.top.height {
					top_score++
					break
				}
				curr = curr.top
				top_score++
			}

			curr = square
			right_score := 0
			for {
				if curr.right == nil {
					visible = true
					break
				}
				if square.height <= curr.right.height {
					right_score++
					break
				}
				curr = curr.right
				right_score++
			}

			curr = square
			bottom_score := 0
			for {
				if curr.bottom == nil {
					visible = true
					break
				}
				if square.height <= curr.bottom.height {
					bottom_score++
					break
				}
				curr = curr.bottom
				bottom_score++
			}

			curr = square
			left_score := 0
			for {
				if curr.left == nil {
					visible = true
					break
				}
				if square.height <= curr.left.height {
					left_score++
					break
				}
				curr = curr.left
				left_score++
			}

			square.scenic_score = top_score * right_score * bottom_score * left_score
			square.visible = visible
		}
	}

	visibles := make([]*Square, 0)
	var most_scenic *Square
	for _, row := range grid {
		for _, square := range row {
			if square.visible {
				visibles = append(visibles, square)
			}
			if most_scenic == nil || most_scenic.scenic_score < square.scenic_score {
				most_scenic = square
			}
		}
	}

	fmt.Println("Part one:", len(visibles))
	fmt.Println("Part two:", most_scenic.scenic_score)
}
