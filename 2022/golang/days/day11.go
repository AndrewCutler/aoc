package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Test struct {
	divisor    int
	true_dest  int
	false_dest int
}

type Monkey struct {
	number              int
	worry_levels        []int // order matters
	transform_as_string string
	test                *Test
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

// have to track how many times this is called
var call_counter map[int]int = make(map[int]int)
func (m *Monkey) inspect() bool {
	call_counter[m.number]++
	if len(m.worry_levels) > 0 {
		boredom_factor := 3
		worry_level := m.worry_levels[0]
		transformed := m.apply_transform(worry_level) / boredom_factor

		return transformed%m.test.divisor == 0
	}
	return false
}

func (m *Monkey) toss(monkeys []*Monkey) {
	is_true := m.inspect()
	worry_level := m.worry_levels[0]
	// is_true := worry_level % m.test.divisor == 0

	for _, curr := range monkeys {
		if (is_true && curr.number == m.test.true_dest) || (!is_true && curr.number == m.test.false_dest) {
			curr.worry_levels = append(curr.worry_levels, worry_level)
			m.worry_levels = m.worry_levels[1:]
			return
		}
	}
}

func (m *Monkey) toss_all(monkeys []*Monkey) {
	for range m.worry_levels {
		m.toss(monkeys)
	}
}

// We have a list of monkeys, now each needs to take a "turn",
// in numeric order, which consists of inspecting all 
// worry_levels until none are left.
// Once every monkey has done this, the "round" is over.

func DayEleven() {
	file, err := os.Open("../data/day11.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var monkeys []*Monkey
	var monkey Monkey
	NEXT_LINE:
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Monkey") {
			number, _ := strconv.Atoi(strings.Replace(strings.Split(line, " ")[1], ":", "", 1))
			monkey = Monkey{
				number: number,
			}
		}

		if strings.HasPrefix(strings.Trim(line, " "), "Starting items: ") {
			items_stringified := strings.Split(line, "Starting items: ")[1]
			item_worry_levels_stringified := strings.Split(items_stringified, ", ")
			for _, worry_level_s := range item_worry_levels_stringified {
				worry_level, _ := strconv.Atoi(worry_level_s)
				monkey.worry_levels = append(monkey.worry_levels, worry_level)
			}
		}

		if strings.HasPrefix(strings.Trim(line, " "), "Operation: ") {
			var operand string
			var operator string
			parsed := strings.Split(strings.Trim(line, " "), " ")
			for i := len(parsed) - 1; i >= 0; i-- {
				if i == len(parsed)-1 {
					operand = parsed[i]
					monkey.transform_as_string += operand
				}

				if i == len(parsed)-2 {
					operator = parsed[i]
					monkey.transform_as_string += operator
				}
			}
		}

		if strings.HasPrefix(strings.Trim(line, " "), "Test: ") {
			parsed := strings.Split(strings.Trim(line, " "), " ")
			divisor, _ := strconv.Atoi(parsed[len(parsed)-1])
			if monkey.test == nil {
				monkey.test = &Test{
					divisor: divisor,
				}
			} else {
				monkey.test.divisor = divisor
			}
		}

		if strings.HasPrefix(strings.Trim(line, " "), "If true: ") {
			parsed := strings.Split(strings.Trim(line, " "), " ")
			true_destination, _ := strconv.Atoi(parsed[len(parsed)-1])
			if monkey.test == nil {
				monkey.test = &Test{
					true_dest: true_destination,
				}
			} else {
				monkey.test.true_dest = true_destination
			}

			for _, curr := range monkeys {
				if curr.number == monkey.number {
					continue NEXT_LINE
				}
			}

			temp := monkey
			monkeys = append(monkeys, &temp)
		}

		if strings.HasPrefix(strings.Trim(line, " "), "If false: ") {
			parsed := strings.Split(strings.Trim(line, " "), " ")
			false_destination, _ := strconv.Atoi(parsed[len(parsed)-1])
			if monkey.test == nil {
				monkey.test = &Test{
					false_dest: false_destination,
				}
			} else {
				monkey.test.false_dest = false_destination
			}

			for _, curr := range monkeys {
				if curr.number == monkey.number {
					continue NEXT_LINE
				}
			}

			temp := monkey
			monkeys = append(monkeys, &temp)
		}
	}
	
	// execute one round
	for _, monkey := range monkeys {
		monkey.toss_all(monkeys)
	}

	// monkeys[0].toss(monkeys)
	// monkeys[0].toss(monkeys)
	fmt.Println(call_counter)
}
