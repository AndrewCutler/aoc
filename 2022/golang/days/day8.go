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
	
	forest := make([][]byte, num_of_cols)
	for i := range forest {
		forest[i] = make([]byte, num_of_rows)
	}
	
	file.Seek(0, io.SeekStart)

	scanner = bufio.NewScanner(file)

	line_num := 0
	zero_byte_offset := 48
	for scanner.Scan() {
		line := scanner.Bytes()

		for i, b := range line {
			forest[line_num][i] = b - byte(zero_byte_offset)
		}

		line_num++
	}

	fmt.Println(forest)
}
