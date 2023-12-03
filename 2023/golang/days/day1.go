package days

import (
	"fmt"
	"os"
)

func DayOne() {
	file, err := os.Open("../data/day1.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fmt.Println("Day one:")
}
