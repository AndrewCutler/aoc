package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	size     int
	name     string
	visited  bool
	parent   *Node
	children []*Node
	// files File -- files here, as struct with name and size
}

func get_sum(root *Node, sum *int) {
	if root == nil || root.visited == true {
		return
	}
	root.visited = true
	if root.size <= 100000 {
		*sum += root.size
	}
	for _, child := range root.children {
		get_sum(child, sum)
	}
	if root.parent != nil {
		get_sum(root.parent, sum)
	}
}

func get_size(root *Node) int {
	if len(root.children) == 0 {
		return root.size
	}

	for _, child := range root.children {
		root.size += get_size(child)
	}

	return root.size
}

func find_closest(root *Node, unused_space int, min int) int {
	if root == nil {
		return min
	}

	if root.size > 30_000_000-unused_space {
		if min == 0 || root.size < min {
			min = root.size
		}
	}
	for _, child := range root.children {
		min = find_closest(child, unused_space, min)
	}

	// this will print the correct answer eventually,
	// but why isn't min the answer at the end?
	fmt.Println(min)
	return min
}

func DaySeven() {
	// if a directory's contents exceed 100,000, add to answer
	part_one := 0

	file, err := os.Open("../data/day7.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var curr *Node
OUTER:
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "\r\n", "")

		if strings.HasPrefix(line, "$ ls") || strings.HasPrefix(line, "dir") {
			continue
		}
		if strings.HasSuffix(line, "cd ..") {
			curr = curr.parent
			continue
		}
		if strings.HasPrefix(line, "$ cd") {
			name := strings.Split(line, " ")[2]

			if curr == nil {
				node := Node{
					name:     name,
					children: make([]*Node, 0),
					visited:  false,
				}
				curr = &node
				continue
			}

			for _, child := range curr.children {
				if child.name == name {
					child.parent = curr
					curr = child
					continue OUTER
				}
			}

			node := Node{
				name:     name,
				children: make([]*Node, 0),
				visited:  false,
				parent:   curr,
			}

			if curr != nil {
				curr.children = append(curr.children, &node)
			}

			curr = &node
			continue
		}

		size, err := strconv.Atoi(strings.Split(line, " ")[0])

		if err != nil {
			panic(err)
		}
		curr.size += size
	}

	// return to root
	for {
		if curr.parent == nil {
			break
		}
		curr = curr.parent
	}

	get_size(curr)

	// now sum all directory sizes <= 100000
	get_sum(curr, &part_one)

	// find directory where 70_000_000 - directory_size > 30_000_000
	// but as close to 30_000_000 as possible

	unused_space := 70_000_000 - curr.size
	part_two := 0
	find_closest(curr, unused_space, part_two)

	fmt.Printf("Part one: %v\n", part_one)
	fmt.Printf("Part two: %v\n", part_two)
}
