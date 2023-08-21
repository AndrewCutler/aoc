package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	number int
	worry_levels []int // order matters
	transform func(int) int // how worry level changes
	toss func(int) // determines where to toss worry_level
}

func DayEleven() {
	fmt.Println("day11")
	file, err := os.Open("../data/day11.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var monkey Monkey
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Monkey") {
			number, _ := strconv.Atoi(strings.Replace(strings.Split(line, " ")[1], ":", "", 1))
			monkey = Monkey{
				number: number,
			}
			// fmt.Printf("monkey: %v", monkey)
			continue
		}

		if strings.HasPrefix(strings.Trim(line, " "), "Starting items: ") {
			items_stringified := strings.Split(line, "Starting items: ")[1]
			item_worry_levels_stringified := strings.Split(items_stringified, ", ")
			for _, worry_level_s := range item_worry_levels_stringified {
				worry_level, _ := strconv.Atoi(worry_level_s)
				monkey.worry_levels = append(monkey.worry_levels, worry_level)
			}
			continue
		}

		fmt.Println(monkey)
	}
}
