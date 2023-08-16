package days

import (
	"bufio"
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
	id      string
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

	zero_byte_offset := 48
	for line_num := 0; scanner.Scan(); line_num++ {
		line := scanner.Bytes()

		// first, allocate squares
		for i, b := range line {
			curr := Square{
				top:     &Square{},
				right:   &Square{},
				bottom:  &Square{},
				left:    &Square{},
				visited: false,
				height:  int(b) - zero_byte_offset,
				id:      string(line_num) + string(i),
			}

			grid[line_num][i] = &curr
		}

		// for i, b := range line {
		// 	curr := grid[line_num][i]
		// 	// why does only top seem to work?
		// 	if line_num != 0 {
		// 		fmt.Printf("top before: %v\n", curr.top)
		// 		curr.top = grid[line_num-1][i]
		// 		fmt.Printf("top after: %v\n", curr.top)
		// 	}
		// 	if line_num != num_of_rows-1 {
		// 		fmt.Printf("bottom before: %v\n", curr.bottom)
		// 		fmt.Printf("bottom square: %v\n", grid[line_num+1][i])
		// 		curr.bottom = grid[line_num+1][i]
		// 		fmt.Printf("bottom after: %v\n", curr.bottom)
		// 	}
		// 	if i != 0 {
		// 		fmt.Printf("left before: %v\n", curr.right)
		// 		curr.left = grid[line_num][i-1]
		// 		fmt.Printf("left after: %v\n", curr.right)
		// 	}
		// 	if i != num_of_cols-1 {
		// 		fmt.Printf("right before: %v\n", curr.right)
		// 		curr.right = grid[line_num][i+1]
		// 		fmt.Printf("right after: %v\n", curr.right)
		// 	}

		// 	curr.height = int(b) - zero_byte_offset

		// 	// grid[line_num][i] = curr
		// }
	}

	for row, _ := range grid {
		for col, _ := range grid[row] {
			curr := grid[row][col]
			// why does only top seem to work?
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

	// visibles := make([]*Square, 0)
	for i, row := range grid {
		if i < 2 {
			// fmt.Println(row)
			for _, square := range row {
				// go up, right, down, left
				// and if you reach nil without encountering
				// any squares with >= height,
				// add to visibles
				temp := square

				// let's go right first
				curr := square
				for {
					// curr.visited = true
					// fmt.Println(curr)
					if curr.right == nil {
						// fmt.Printf("rightmost: %v\n", curr)
						curr = temp
						break
					}
					curr = curr.right
				}
			}
		}
	}
}
