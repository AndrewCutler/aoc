package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	number              int
	worry_levels        []int         // order matters
	transform_as_string string
	toss                func(int) // determines where to toss worry_level
}

func (m *Monkey) apply_transform(i int) int {
	t := m.transform_as_string
	operand := t[:len(t)-1]
	operator := t[len(t)-1:]

	if operand == "old" {
		switch operator {
		case "+":
			{
				return i + i
			}
		case "*":
			{
				return i * i
			}
		}
	}

	operand_value, _ := strconv.Atoi(operand)
	switch operator {
	case "+":
		{
			return operand_value + i
		}
	case "*":
		{
			return operand_value * i
		}
	}

	return 0
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

		if strings.HasPrefix(strings.Trim(line, " "), "Operation: ") {
			var operand string
			var operator string
			operation := strings.Split(strings.Trim(line, " "), " ")
			for i := len(operation) - 1; i >= 0; i-- {
				if i == len(operation)-1 {
					operand = operation[i]
					monkey.transform_as_string += operand
				}

				if i == len(operation)-2 {
					operator = operation[i]
					monkey.transform_as_string += operator
				}
			}
		}

		fmt.Println(monkey)
	}
}
