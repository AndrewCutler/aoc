package days

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Square struct {
	top     *Square
	right   *Square
	bottom  *Square
	left    *Square
	visited bool
	height  int
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
	file, err := os.Open("../data/day8.test.txt")

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

	line_num := 0
	zero_byte_offset := 48
	for scanner.Scan() {
		line := scanner.Bytes()

		for i, b := range line {
			curr := new(Square)
			if line_num != 0 {
				curr.top = grid[line_num-1][i]
			}
			if line_num != num_of_rows-1 {
				curr.bottom = grid[line_num+1][i]
			}
			if i != 0 {
				curr.left = grid[line_num][i-1]
			}
			if i != num_of_cols-1 {
				curr.right = grid[line_num][i+1]
			}

			curr.height = int(b) - zero_byte_offset

			grid[line_num][i] = curr
		}

		line_num++
	}

	// visibles := make([]*Square, 0)
	for i, row := range grid {
		if i < 2 {
			// fmt.Println(row)
			for j, _ := range row {
				curr := *grid[i][j]
				fmt.Println(curr)
				fmt.Println(*grid[i][j])
				// fmt.Println(col)
				// go up, right, down, left
				// and if you reach nil without encountering
				// any squares with >= height,
				// add to visibles
			}
		}
	}
}
