package days

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Square struct {
	top         *Square
	right       *Square
	bottom      *Square
	left        *Square
	visited     bool
	tree_height int
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

	forest := make([][]Square, num_of_cols)
	for i := range forest {
		forest[i] = make([]Square, num_of_rows)
	}

	file.Seek(0, io.SeekStart)

	scanner = bufio.NewScanner(file)

	line_num := 0
	zero_byte_offset := 48
	for scanner.Scan() {
		line := scanner.Bytes()

		for i, b := range line {
			tree := new(Square)
			if line_num != 0 {
				tree.top = &forest[line_num-1][i]
			}
			if line_num != num_of_rows-1 { // todo: verify -1 is correct
				tree.bottom = &forest[line_num+1][i]
			}
			if i != 0 {
				tree.left = &forest[line_num][i-1]
			}
			if i != num_of_cols-1 { // todo: verify -1 is correct
				tree.right = &forest[line_num][i+1]
			}

			tree.tree_height = int(b) - zero_byte_offset

			forest[line_num][i] = *tree
		}

		line_num++
	}

	fmt.Println(forest)
}
