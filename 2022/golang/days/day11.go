package days

import (
	"bufio"
	"fmt"
	"os"
)

func DayEleven() {
	fmt.Println("day11")
	file, err := os.Open("../data/day11.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
